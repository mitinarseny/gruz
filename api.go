package gruz

import (
    "context"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strconv"
    "strings"
    "time"
)

type Language int

type PersonType string

const (
    baseURL          = "http://ruz.hse.ru/api"
    scheduleEndpoint = "schedule"
    timeZone = "Europe/Moscow"

    RussianLanguage Language = 1
    EnglishLanguage Language = 2

    StudentPerson  PersonType = "student"
    LecturerPerson PersonType = "lecturer"
)

type APIError struct {
    statusCode int
    msg        string
}

func NewAPIError(statusCode int, msg string) *APIError {
    return &APIError{
        statusCode: statusCode,
        msg:        msg,
    }
}

func (err APIError) Error() string {
    return fmt.Sprintf("%d: %v", err.statusCode, err.msg)
}

type Client struct {
    httpClient *http.Client
}

func NewClient(client *http.Client) *Client {
    return &Client{
        httpClient: client,
    }
}

func (c *Client) GetSchedule(ctx context.Context,
    hseID int64, personType PersonType,
    fromDate time.Time, toDate time.Time, lang Language, ) ([]Lesson, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", strings.Join([]string{
        baseURL,
        scheduleEndpoint,
        string(personType),
        strconv.FormatInt(hseID, 10),
    }, "/"), nil)
    if err != nil {
        return nil, err
    }

    q := req.URL.Query()
    q.Set("start", ruzDate(&fromDate))
    q.Set("finish", ruzDate(&toDate))
    q.Set("lng", strconv.Itoa(int(lang)))
    req.URL.RawQuery = q.Encode()

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if !(200 <= resp.StatusCode && resp.StatusCode < 300) {
        b, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return nil, err
        }
        return nil, NewAPIError(resp.StatusCode, string(b))
    }

    res := make([]Lesson, 0)
    if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
        return nil, err
    }
    return res, nil
}
