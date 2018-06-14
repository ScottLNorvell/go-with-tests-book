package main

import (
  "fmt"
  "encoding/json"
  "io"
  "net/http"
  "net/http/httptest"
  "reflect"
  "testing"
)

type StubPlayerStore struct {
  stores map[string]int
  winCalls []string
  league []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
  score := s.stores[name]
  return score
}

func (s *StubPlayerStore) RecordWin(name string) {
  s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
  return s.league
}

func TestGETPlayers(t *testing.T) {
  store := StubPlayerStore{
    map[string]int{
      "Pepper": 20,
      "Floyd": 10,
    },
    nil,
    nil,
  }
  server := NewPlayerServer(&store)

  t.Run("returns Pepper's score", func (t *testing.T) {
    req := newGetScoreRequest("Pepper")
    resp := httptest.NewRecorder()

    server.ServeHTTP(resp, req)

    assertStatus(t, resp.Code, http.StatusOK)
    assertResponseBody(t, resp.Body.String(), "20")
  })

  t.Run("returns Floyd's score", func (t *testing.T) {
    req := newGetScoreRequest("Floyd")
    resp := httptest.NewRecorder()

    server.ServeHTTP(resp, req)

    assertStatus(t, resp.Code, http.StatusOK)
    assertResponseBody(t, resp.Body.String(), "10")
  })

  t.Run("returns 404 for missing players", func (t *testing.T) {
    req := newGetScoreRequest("Jerry")
    resp := httptest.NewRecorder()

    server.ServeHTTP(resp, req)

    got := resp.Code
    want := http.StatusNotFound

    assertStatus(t, got, want)
  })
}

func TestStoreWins(t *testing.T) {
  store := StubPlayerStore{
    map[string]int{
      "Pepper": 20,
      "Floyd": 10,
    },
    nil,
    nil,
  }
  server := NewPlayerServer(&store)

  t.Run("it records wins on a POST", func (t *testing.T) {
    player := "Pepper"

    req := newPostWinRequest(player)
    resp := httptest.NewRecorder()

    server.ServeHTTP(resp, req)

    assertStatus(t, resp.Code, http.StatusAccepted)

    if len(store.winCalls) != 1 {
      t.Errorf("got %d calls to RecordWin; want %d", len(store.winCalls), 1)
    }

    if store.winCalls[0] != player {
      t.Errorf("did not call RecordWin with correct player got '%s', want '%s'", store.winCalls[0], player)
    }
  })
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
  store := NewInMemoryPlayerStore()
  server := NewPlayerServer(store)
  player := "Pepper"

  server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
  server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
  server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

  t.Run("get score", func (t *testing.T) {
    resp := httptest.NewRecorder()
    server.ServeHTTP(resp, newGetScoreRequest(player))

    assertStatus(t, resp.Code, http.StatusOK)

    assertResponseBody(t, resp.Body.String(), "3")
  })

  t.Run("get league", func (t *testing.T) {
    resp := httptest.NewRecorder()
    server.ServeHTTP(resp, newGetLeagueRequest())

    got := getLeagueFromResponse(t, resp.Body)

    want := []Player{
      {"Pepper", 3},
    }

    assertLeague(t, got, want)
  })
}

func TestLeague(t *testing.T) {
  store := StubPlayerStore{}
  server := NewPlayerServer(&store)

  t.Run("it returns 200 for /league", func (t *testing.T) {
    req, _ := http.NewRequest(http.MethodGet, "/league", nil)
    resp := httptest.NewRecorder()

    server.ServeHTTP(resp, req)

    var got []Player

    err := json.NewDecoder(resp.Body).Decode(&got)

    if err != nil {
      t.Fatalf("unable to parse response from server '%s' into slice of Player '%v'", resp.Body, err)
    }

    assertStatus(t, resp.Code, http.StatusOK)
  })

  t.Run("it returns a league table of JSON", func (t *testing.T) {
    wantedLeague := []Player{
      {"Scott", 80},
      {"Chris", 90},
      {"Tamara", 70},
    }

    store := StubPlayerStore{nil, nil, wantedLeague}

    server := NewPlayerServer(&store)

    req := newGetLeagueRequest()
    resp := httptest.NewRecorder()

    server.ServeHTTP(resp, req)

    got := getLeagueFromResponse(t, resp.Body)

    assertStatus(t, resp.Code, http.StatusOK)

    assertContentType(t, resp, jsonContentType)

    assertLeague(t, got, wantedLeague)
  })
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
  t.Helper()

  err := json.NewDecoder(body).Decode(&league)

  if err != nil {
    t.Fatalf("unable to parse response from server '%s' into slice of Player '%v'", body, err)
  }

  return
}

func newGetLeagueRequest() *http.Request {
  req, _ := http.NewRequest(http.MethodGet, "/league", nil)
  return req
}

func newPostWinRequest(name string) *http.Request {
  req, _ := http.NewRequest(
    http.MethodPost,
    fmt.Sprintf("/players/%s", name),
    nil,
  )
  return req
}

func newGetScoreRequest(name string) *http.Request {
  req, _ := http.NewRequest(
    http.MethodGet,
    fmt.Sprintf("/players/%s", name),
    nil,
  )
  return req
}

func assertStatus(t *testing.T, got, want int) {
  t.Helper()
  if got != want {
    t.Errorf("got status %d want %d", got, want)
  }
}

func assertResponseBody(t *testing.T, got, want string) {
  t.Helper()
  if got != want {
    t.Errorf("response body incorrect: got '%s' want '%s'", got, want)
  }
}

func assertLeague(t *testing.T, got, want []Player) {
  if !reflect.DeepEqual(got, want) {
    t.Errorf("got %v want %v", got, want)
  }
}
const jsonContentType = "application/json"
func assertContentType(t *testing.T, resp *httptest.ResponseRecorder, want string) {
  t.Helper()

  if resp.Header().Get("content-type") != want {
    t.Errorf("response didn't have content-type of 'application/json' got %v", resp.HeaderMap)
  }
}
