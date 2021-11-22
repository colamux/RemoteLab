package raspi

import (
	"bytes"
	"fmt"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

const (
	// START_NGINX = "/usr/local/nginx/sbin/nginx -c "
	FFMPEG_WITH_VIDEO0 = "ffmpeg -i /dev/video0"
	FFMPEG_WITH_VIDEO1 = ""
	FFMPEG_WITH_VIDEO2 = ""
	FFMPEG_WITH_VIDEO3 = ""
)

type Connect struct {
	client *ssh.Client
}

type Session struct {
	Std1    *bytes.Buffer
	Std2    *bytes.Buffer
	Session *ssh.Session
}

func New(user, password, host string, port int) (*Connect, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		err          error
	)

	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: func(host string, remote net.Addr, key ssh.PublicKey) error { return nil },
	}

	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	return &Connect{client: client}, nil
}

func (r *Connect) Close() {
	r.client.Close()
	r.client = nil
}

func (r *Connect) CreateSession() (*Session, error) {

	// var session *ssh.Session
	session, err := r.client.NewSession()
	if err != nil {
		return nil, err
	}

	// byt

	return &Session{Session: session, Std1: &bytes.Buffer{}, Std2: &bytes.Buffer{}}, nil
}

func (s *Session) Run(cmd string) error {
	s.Session.Stdout = s.Std1
	s.Session.Stderr = s.Std2
	if err := s.Session.Run(cmd); err != nil {
		return err
	}
	return nil
}

func (s *Session) Close() {
	s.Session.Close()
	s.Session = nil
	s.Std1 = nil
	s.Std2 = nil
}
