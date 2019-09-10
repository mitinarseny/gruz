package gruz

import (
    "encoding/json"
    "strings"
    "time"
)

type Lesson struct {
    Auditorium         string
    AuditoriumAmount   int
    AuditoriumOid      int
    Building           string
    CreatedAt          time.Time
    DetailInfo         string
    Discipline         string
    DisciplineOid      int
    DisciplineInPlan   string
    DisciplineTypeLoad int
    End                time.Time
    Group              interface{}
    GroupOid           int
    HideInCapacity     int
    IsBan              bool
    KindOfWork         string
    Lecturer           string
    LecturerOid        int
    LecturerUID        string
    ModifiedAt         time.Time
    Start              time.Time
    SubGroup           interface{}
    SubGroupOid        int
}

type jsonLesson struct {
    Auditorium         string      `json:"auditorium"`
    AuditoriumAmount   int         `json:"auditoriumAmount"`
    AuditoriumOid      int         `json:"auditoriumOid"`
    BeginLesson        string      `json:"beginLesson"`
    Building           string      `json:"building"`
    CreatedDate        string      `json:"createddate"`
    Date               string      `json:"date"`
    DetailInfo         string      `json:"detailInfo"`
    Discipline         string      `json:"discipline"`
    DisciplineOid      int         `json:"disciplineOid"`
    DisciplineInPlan   string      `json:"disciplineinplan"`
    DisciplineTypeLoad int         `json:"disciplinetypeload"`
    EndLesson          string      `json:"endLesson"`
    Group              interface{} `json:"group"`
    GroupOid           int         `json:"groupOid"`
    HideInCapacity     int         `json:"hideincapacity"`
    IsBan              bool        `json:"isBan"`
    KindOfWork         string      `json:"kindOfWork"`
    Lecturer           string      `json:"lecturer"`
    LecturerOid        int         `json:"lecturerOid"`
    LecturerUID        string      `json:"lecturerUID"`
    ModifiedDate       string      `json:"modifieddate"`
    SubGroup           interface{} `json:"subGroup"`
    SubGroupOid        int         `json:"subGroupOid"`
}

func (l *Lesson) UnmarshalJSON(b []byte) error {
    var aux jsonLesson
    if err := json.Unmarshal(b, &aux); err != nil {
        return err
    }

    createdAt, err := time.Parse(time.RFC3339, strings.TrimSuffix(aux.CreatedDate, "00:00"))
    if err != nil {
        createdAt = time.Now() // treat as it has been just created
    }
    l.CreatedAt = createdAt

    modifiedAt, err := time.Parse(time.RFC3339, strings.TrimSuffix(aux.ModifiedDate, "00:00"))
    if err != nil {
        modifiedAt = l.CreatedAt // should have named package "balast" instead of "gruz"
    }
    l.ModifiedAt = modifiedAt

    start, err := time.Parse("2006.01.02T15:04Z07:00",
        aux.Date+"T"+aux.BeginLesson+"+03:00")
    if err != nil {
        return err
    }
    l.Start = start

    end, err := time.Parse("2006.01.02T15:04Z07:00",
        aux.Date+"T"+aux.EndLesson+"+03:00")
    if err != nil {
        return err
    }
    l.End = end

    l.Auditorium = aux.Auditorium
    l.AuditoriumAmount = aux.AuditoriumAmount
    l.AuditoriumOid = aux.AuditoriumOid
    l.Building = aux.Building
    l.DetailInfo = aux.DetailInfo
    l.Discipline = aux.Discipline
    l.DisciplineOid = aux.DisciplineOid
    l.DisciplineInPlan = aux.DisciplineInPlan
    l.DisciplineTypeLoad = aux.DisciplineTypeLoad
    l.Group = aux.Group
    l.GroupOid = aux.GroupOid
    l.HideInCapacity = aux.HideInCapacity
    l.IsBan = aux.IsBan
    l.KindOfWork = aux.KindOfWork
    l.Lecturer = aux.Lecturer
    l.LecturerOid = aux.LecturerOid
    l.LecturerUID = aux.LecturerUID
    l.SubGroup = aux.SubGroup
    l.SubGroupOid = aux.SubGroupOid

    return nil
}
