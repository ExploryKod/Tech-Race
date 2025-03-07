package services

import (
	"fmt"
	"hetic/tech-race/internal/models"
	"hetic/tech-race/internal/mqtt"
	"time"
)

type SessionServiceID struct {
	ID int
}

type SessionService struct {
	db models.DatabaseInterface
}

func NewSessionService(db models.DatabaseInterface) *SessionService {
	return &SessionService{db: db}
}

func (s *SessionService) GetCurrentSessionID() (SessionServiceID, error) {
	sessionID, err := s.db.GetCurrentSessionID()
	if err != nil {
		return SessionServiceID{}, err
	}
	return SessionServiceID{ID: sessionID}, nil
}

func (s *SessionService) GetAll() ([]models.Session, error) {
	sessions, err := s.db.GetAll()
	if err != nil {
		return nil, err
	}
	fmt.Println(sessions)
	return sessions, nil
}
func (s *SessionService) Start(isAutopilot bool) error {
	timeStamp := time.Now()
	err := s.db.StartSession(timeStamp, isAutopilot)
	if err != nil {
		return err
	}
	mqttClient := mqtt.NewMQTTClient(s.db)
	_ = mqttClient.ConnectAndSubscribe(isAutopilot)
	fmt.Println("Session started")
	return nil
}

func (s *SessionService) Stop() error {
	timeStamp := time.Now()
	err := s.db.StopSession(timeStamp)
	if err != nil {
		return err
	}
	mqttClient := mqtt.NewMQTTClient(s.db)
	mqttClient.Disconnect()
	fmt.Println("Session stopped")
	return nil
}
func (s *SessionService) IsSessionActive() (bool, error) {
	return s.db.IsSessionActive()
}

func (s *SessionService) GetAllSessionInfo() ([]models.SessionInfo, error) {
	sessions, err := s.db.GetAll()
	if err != nil {
		return nil, err
	}

	var sessionInfos []models.SessionInfo
	for _, session := range sessions {
		startDate := session.StartDate.Format("02/01/2006 - 15:04:05")
		durationInSeconds := int64(session.EndDate.Sub(session.StartDate).Seconds())
		endDate := session.EndDate.Format("02/01/2006 - 15:04:05")
		duration := session.EndDate.Sub(session.StartDate)
		hours := int(duration.Hours())
		minutes := int(duration.Minutes()) % 60
		seconds := int(duration.Seconds()) % 60
		//durationStr := fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
		durationStr := ""
		if hours > 0 {
			durationStr = fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
		}
		if minutes > 0 {
			durationStr = fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
		} else {
			durationStr = fmt.Sprintf("%ds", seconds)
		}

		collisions, err := s.db.GetCollisionsBySessionID(session.ID)
		if err != nil {
			return nil, err
		}

		tracks, err := s.db.GetTracksBySessionID(session.ID)
		if err != nil {
			return nil, err
		}
		videos, err := s.db.GetVideosBySessionID(session.ID)
		if err != nil {
			return nil, err
		}

		videoInfo := models.VideoInfo{
			VideoURLs: make([]string, len(videos)),
		}

		for i, video := range videos {
			videoInfo.VideoURLs[i] = video.VideoURL
		}
		collisionInfo := models.CollisionInfo{
			Count:      len(collisions),
			Distances:  make([]float64, len(collisions)),
			Timestamps: make([]string, len(collisions)),
		}
		for i, collision := range collisions {
			timingCollision := collision.Timestamp.Sub(session.StartDate)
			hours := int(timingCollision.Hours())
			minutes := int(timingCollision.Minutes()) % 60
			seconds := int(timingCollision.Seconds()) % 60
			//collisionInfo.Timestamps[i] = fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
			collisionInfo.Timestamps[i] = ""
			if hours > 0 {
				collisionInfo.Timestamps[i] = fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
			}
			if minutes > 0 {
				collisionInfo.Timestamps[i] = fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
			} else {
				collisionInfo.Timestamps[i] = fmt.Sprintf("%ds", seconds)
			}
			collisionInfo.Distances[i] = collision.Distance
		}

		trackInfo := models.TrackInfo{
			Count:              len(tracks),
			LineTrackingValues: make([]int, len(tracks)),
			Timestamps:         make([]string, len(tracks)),
		}
		for i, track := range tracks {
			timingTrack := track.Timestamp.Sub(session.StartDate)
			hours := int(timingTrack.Hours())
			minutes := int(timingTrack.Minutes()) % 60
			seconds := int(timingTrack.Seconds()) % 60
			//trackInfo.Timestamps[i] = fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
			trackInfo.Timestamps[i] = ""
			if hours > 0 {
				trackInfo.Timestamps[i] = fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
			}
			if minutes > 0 {
				trackInfo.Timestamps[i] = fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
			} else {
				trackInfo.Timestamps[i] = fmt.Sprintf("%ds", seconds)

			}

			trackInfo.LineTrackingValues[i] = track.LineTrackingValue
		}

		sessionInfo := models.SessionInfo{
			ID:                session.ID,
			StartDate:         startDate,
			EndDate:           endDate,
			DurationInSeconds: durationInSeconds,
			Duration:          durationStr,
			IsAutopilot:       session.IsAutopilot,
			Videos:            videoInfo,
			Collisions:        []models.CollisionInfo{collisionInfo},
			Tracks:            []models.TrackInfo{trackInfo},
		}

		sessionInfos = append(sessionInfos, sessionInfo)
	}

	return sessionInfos, nil
}
