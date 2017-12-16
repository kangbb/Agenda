package entity

import "strings"

// Meeting .
type Meeting struct {
	Sponsor            string
	Participators      []string
	StartDate, EndDate Date
	Title             string
}

func (m_meeting Meeting) init(t_Sponsor string, t_Participators []string, t_StartDate Date, t_EndDate Date, t_Title string) {
	m_meeting.Sponsor = t_Sponsor
	m_meeting.SetParticipator(t_Participators)
	m_meeting.StartDate.CopyDate(t_StartDate)
	m_meeting.EndDate.CopyDate(t_EndDate)
	m_meeting.Title = t_Title
}
func (m_meeting Meeting) CopyMeeting(t_meeting Meeting) {
	m_meeting.Sponsor = t_meeting.Sponsor
	m_meeting.SetParticipator(t_meeting.Participators)
	m_meeting.StartDate.CopyDate(t_meeting.StartDate)
	m_meeting.EndDate.CopyDate(t_meeting.EndDate)
	m_meeting.Title = t_meeting.Title
}
func (m_meeting Meeting) GetSponsor() string {
	return m_meeting.Sponsor
}

/**
* @brief set the sponsor of a meeting
* @param  the new sponsor string
 */
func (m_meeting Meeting) SetSponsor(t_sponsor string) {
	m_meeting.Sponsor = t_sponsor
}

/**
* @brief  get the participator of a meeting
* @return return a string indicate participator
 */
func (m_meeting Meeting) GetParticipator() []string {
	return m_meeting.Participators
}

/**
*   @brief set the new participator of a meeting
*   @param the new participator string
 */

func (m_meeting Meeting) SetParticipator(t_participators []string) {
	var length = len(t_participators)
	for i := 0; i < length; i++ {
		m_meeting.Participators[i] = t_participators[i]
	}
}

/**
* @brief get the startDate of a meeting
* @return return a string indicate startDate
 */
func (m_meeting Meeting) GetStartDate() Date {
	return m_meeting.StartDate
}

/**
* @brief  set the startDate of a meeting
* @param  the new startdate of a meeting
 */
func (m_meeting Meeting) SetStartDate(t_startTime Date) {
	m_meeting.StartDate.CopyDate(t_startTime)
}

func (m_meeting Meeting) GetEndDate() Date {
	return m_meeting.EndDate
}

func (m_meeting Meeting) SetEndDate(t_endTime Date) {
	m_meeting.EndDate.CopyDate(t_endTime)
}

func (m_meeting Meeting) GetTitle() string {
	return m_meeting.Title
}

func (m_meeting Meeting) SetTitle(t_title string) {
	m_meeting.Title = t_title
}

func (m_meeting Meeting) IsParticipator(t_username string) bool {
	for i := 0; i < len(m_meeting.Participators); i++ {
		if strings.EqualFold(m_meeting.Participators[i], t_username) == true {
			return true
		}
	}
	return false
}
