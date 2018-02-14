//
// Copyright 2015, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package gitlab

import (
	"fmt"
	"net/url"

	"github.com/xanzy/go-gitlab"
)

// BoardsService handles communication with the issue related methods
// of the GitLab API.
//
// GitLab API docs:
// https://gitlab.com/gitlab-org/gitlab-ce/blob/8-16-stable/doc/api/boards.md
type BoardsService struct {
	client *Client
}

// Board represents a GitLab board.
//
// GitLab API docs:
// https://gitlab.com/gitlab-org/gitlab-ce/blob/8-16-stable/doc/api/boards.md
type Board struct {
	ID    int         `json:"id"`
	Lists []BoardList `json:"lists"`
}

// BoardList represents a GitLab board list item.
//
// GitLab API docs:
// https://gitlab.com/gitlab-org/gitlab-ce/blob/8-16-stable/doc/api/boards.md
type BoardList struct {
	ID       int          `json:"id"`
	Label    gitlab.Label `json:"label"`
	Position int          `json:"position"`
}

// ListProjectBoards gets a list of all boards of project.
//
// GitLab API docs:
// https://gitlab.com/gitlab-org/gitlab-ce/blob/8-16-stable/doc/api/boards.md#project-board
func (*BoardsService) ListProjectBoards(pid interface{}) ([]Board, *Response, error) {
	project, err := parseID(pid)
	if err != nil {
		return nil, nil, err
	}
	u := fmt.Sprintf("projects/%s/boards", url.QueryEscape(project))

	req, err := s.client.NewRequest("GET", u, opt, options)
	if err != nil {
		return nil, nil, err
	}

	var i []*Issue
	resp, err := s.client.Do(req, &i)
	if err != nil {
		return nil, resp, err
	}

	return i, resp, err
}
