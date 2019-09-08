package gruz

import (
    "context"
    "crypto/tls"
    "encoding/json"
    "fmt"
    "net/http"
    "testing"
    "time"

    "github.com/stretchr/testify/require"
)

func TestLesson_UnmarshalJSON(t *testing.T) {
    r := require.New(t)

    rawJSON := `
{
    "auditorium": "101",
    "auditoriumAmount": 92,
    "auditoriumOid": 3269,
    "author": "STAFF\\egrishina",
    "beginLesson": "09:00",
    "building": "Мясницкая ул., д. 20",
    "createddate": "2019-06-26T16:18:41Z00:00",
    "date": "2019.09.03",
    "dateOfNest": "\/Date(1567458000000+0300)\/",
    "dayOfWeek": 2,
    "dayOfWeekString": "Вт",
    "detailInfo": "",
    "discipline": "Социологическая теория (рус)",
    "disciplineOid": 52792,
    "disciplineinplan": "2043984560",
    "disciplinetypeload": 7,
    "endLesson": "10:20",
    "group": null,
    "groupOid": 0,
    "hideincapacity": 0,
    "isBan": false,
    "kindOfWork": "Лекция",
    "lecturer": "доц. Николаев Владимир Геннадьевич",
    "lecturerOid": 34354,
    "lecturerUID": "555711875",
    "lessonNumberEnd": 1,
    "lessonNumberStart": 1,
    "modifieddate": "2019-06-26T16:18:41Z00:00",
    "parentschedule": "2019\/2020_2 курс_1 семестр_Расписание занятий  ОП \"Социология\"",
    "stream": "Ст_Б2018_СОЦЛ:1#П#Социологическая теория#БСЦ181-БСЦ182-БСЦ183-БСЦ184-БСЦ185-БСЦ186",
    "streamOid": 106866,
    "subGroup": null,
    "subGroupOid": 0
}
`
    var l Lesson
    err := json.Unmarshal([]byte(rawJSON), &l)
    r.NoError(err, "There should be no error while unmarshaling raw json string")

    expected := Lesson{
        Auditorium:         "101",
        AuditoriumAmount:   92,
        AuditoriumOid:      3269,
        Building:           "Мясницкая ул., д. 20",
        CreatedAt:          time.Date(2019, 6, 26, 16, 18, 41, 0, time.UTC),
        DetailInfo:         "",
        Discipline:         "Социологическая теория (рус)",
        DisciplineOid:      52792,
        DisciplineInPlan:   "2043984560",
        DisciplineTypeLoad: 7,
        End:                time.Date(2019, 9, 3, 10, 20, 0, 0, nil),
        Group:              nil,
        GroupOid:           0,
        HideInCapacity:     0,
        IsBan:              false,
        KindOfWork:         "Лекция",
        Lecturer:           "доц. Николаев Владимир Геннадьевич",
        LecturerOid:        34354,
        LecturerUID:        "555711875",
        ModifiedAt:         time.Date(2019, 6, 26, 16, 18, 41, 0, time.UTC),
        Start:              time.Date(2019, 9, 3, 9, 0, 0, 0, nil),
        SubGroup:           nil,
        SubGroupOid:        0,
    }
    r.Equal(expected, l)
}

func TestClient_GetSchedule(t *testing.T) {
    // r := require.New(t)
    cl := NewClient(&http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true, // TODO: remove insecure transport
            },
        },
    })
    fmt.Println(cl.GetSchedule(context.Background(),
        173224,
        "student",
        time.Now(),
        time.Now().AddDate(0, 1, 0),
        RussianLanguage))
}
