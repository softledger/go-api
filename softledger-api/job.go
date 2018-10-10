package softledger

import (
	"context"
	"fmt"
)

type JobService service

type Job struct {
	ID          *int64  `json:"_id"`
	Number      *string `json:"number"`
	Name        *string `json:"name"`
	Status      *string `json:"status"`
	Description *string `json:"description"`
	AgentId     *int64  `json:"AgentId"`
	Agent       *Customer
}

type jobResponse struct {
	Data      []*Job `json:"data"`
	TotalJobs int    `json:"totalJobs"`
}

func (cc Job) String() string {
	return Stringify(cc)
}

func (s *JobService) All(ctx context.Context, qry *QueryParams) ([]*Job, int, *Response, error) {

	u, err := addParams("/jobs", qry)

	if err != nil {
		return nil, 0, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, 0, nil, err
	}

	var ccs *jobResponse
	resp, err := s.client.Do(ctx, req, &ccs)
	if err != nil {
		return nil, 0, resp, err
	}

	return ccs.Data, ccs.TotalJobs, resp, nil

}

func (s *JobService) One(ctx context.Context, _id int64) (*Job, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/jobs", _id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cc *Job
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil
}

func (s *JobService) Create(ctx context.Context, payload *Job) (*Job, *Response, error) {

	u := fmt.Sprintf("/jobs")

	req, err := s.client.NewRequest("POST", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Job
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *JobService) Update(ctx context.Context, _id int64, payload *Job) (*Job, *Response, error) {

	u := fmt.Sprintf("%v/%v", "/jobs", _id)

	req, err := s.client.NewRequest("PUT", u, payload)
	if err != nil {
		return nil, nil, err
	}

	var cc *Job
	resp, err := s.client.Do(ctx, req, &cc)
	if err != nil {
		return nil, resp, err
	}

	return cc, resp, nil

}

func (s *JobService) Delete(ctx context.Context, _id int64) (*Response, error) {

	u := fmt.Sprintf("%v/%v", "/jobs", _id)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil

}
