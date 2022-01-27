// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.5

package storageV1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	v1 "github.com/realotz/mstore/api/core/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type VolumeServiceHTTPServer interface {
	CreateFile(context.Context, *CreateFileReq) (*v1.Empty, error)
	CreateVolume(context.Context, *CreateVolumeReq) (*Volume, error)
	DelFile(context.Context, *DelFileReq) (*v1.Empty, error)
	DeleteVolume(context.Context, *DeleteVolumeReq) (*v1.Empty, error)
	FileData(context.Context, *FileReq) (*FileDataRes, error)
	FileDown(context.Context, *FileReq) (*FileDownRes, error)
	ListFile(context.Context, *ListFileReq) (*ListFileReply, error)
	ListVolume(context.Context, *ListVolumeReq) (*ListVolumeReply, error)
	MoveAndCopyFile(context.Context, *MoveCopyFileReq) (*v1.Empty, error)
	RenameFile(context.Context, *RenameFileReq) (*v1.Empty, error)
}

func RegisterVolumeServiceHTTPServer(s *http.Server, srv VolumeServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/v1/volume", _VolumeService_CreateVolume0_HTTP_Handler(srv))
	r.DELETE("/api/v1/volume/{id}", _VolumeService_DeleteVolume0_HTTP_Handler(srv))
	r.GET("/api/v1/volume", _VolumeService_ListVolume0_HTTP_Handler(srv))
	r.GET("/api/v1/volume/{id}/files", _VolumeService_ListFile0_HTTP_Handler(srv))
	r.POST("/api/v1/volume/{id}/files", _VolumeService_CreateFile0_HTTP_Handler(srv))
	r.POST("/api/v1/volume/files/del", _VolumeService_DelFile0_HTTP_Handler(srv))
	r.POST("/api/v1/volume/files/copy-move", _VolumeService_MoveAndCopyFile0_HTTP_Handler(srv))
	r.POST("/api/v1/volume/{id}/files/rename", _VolumeService_RenameFile0_HTTP_Handler(srv))
	r.GET("/api/v1/volume/{id}/files/down", _VolumeService_FileDown0_HTTP_Handler(srv))
	r.GET("/api/v1/files/{id}", _VolumeService_FileData0_HTTP_Handler(srv))
}

func _VolumeService_CreateVolume0_HTTP_Handler(srv VolumeServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateVolumeReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.service.storage.v1.volume.VolumeService/CreateVolume")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateVolume(ctx, req.(*CreateVolumeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Volume)
		return ctx.Result(200, reply)
	}
}

func _VolumeService_DeleteVolume0_HTTP_Handler(srv VolumeServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteVolumeReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.service.storage.v1.volume.VolumeService/DeleteVolume")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteVolume(ctx, req.(*DeleteVolumeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Empty)
		return ctx.Result(200, reply)
	}
}

func _VolumeService_ListVolume0_HTTP_Handler(srv VolumeServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListVolumeReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.service.storage.v1.volume.VolumeService/ListVolume")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListVolume(ctx, req.(*ListVolumeReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListVolumeReply)
		return ctx.Result(200, reply)
	}
}

func _VolumeService_ListFile0_HTTP_Handler(srv VolumeServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListFileReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.service.storage.v1.volume.VolumeService/ListFile")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListFile(ctx, req.(*ListFileReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListFileReply)
		return ctx.Result(200, reply)
	}
}

func _VolumeService_CreateFile0_HTTP_Handler(srv VolumeServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateFileReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.service.storage.v1.volume.VolumeService/CreateFile")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateFile(ctx, req.(*CreateFileReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Empty)
		return ctx.Result(200, reply)
	}
}

func _VolumeService_DelFile0_HTTP_Handler(srv VolumeServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DelFileReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.service.storage.v1.volume.VolumeService/DelFile")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DelFile(ctx, req.(*DelFileReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Empty)
		return ctx.Result(200, reply)
	}
}

func _VolumeService_MoveAndCopyFile0_HTTP_Handler(srv VolumeServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MoveCopyFileReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.service.storage.v1.volume.VolumeService/MoveAndCopyFile")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MoveAndCopyFile(ctx, req.(*MoveCopyFileReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Empty)
		return ctx.Result(200, reply)
	}
}

func _VolumeService_RenameFile0_HTTP_Handler(srv VolumeServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RenameFileReq
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.service.storage.v1.volume.VolumeService/RenameFile")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RenameFile(ctx, req.(*RenameFileReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v1.Empty)
		return ctx.Result(200, reply)
	}
}

func _VolumeService_FileDown0_HTTP_Handler(srv VolumeServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FileReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.service.storage.v1.volume.VolumeService/FileDown")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FileDown(ctx, req.(*FileReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FileDownRes)
		return ctx.Result(200, reply)
	}
}

func _VolumeService_FileData0_HTTP_Handler(srv VolumeServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FileReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.service.storage.v1.volume.VolumeService/FileData")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FileData(ctx, req.(*FileReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FileDataRes)
		return ctx.Result(200, reply)
	}
}

type VolumeServiceHTTPClient interface {
	CreateFile(ctx context.Context, req *CreateFileReq, opts ...http.CallOption) (rsp *v1.Empty, err error)
	CreateVolume(ctx context.Context, req *CreateVolumeReq, opts ...http.CallOption) (rsp *Volume, err error)
	DelFile(ctx context.Context, req *DelFileReq, opts ...http.CallOption) (rsp *v1.Empty, err error)
	DeleteVolume(ctx context.Context, req *DeleteVolumeReq, opts ...http.CallOption) (rsp *v1.Empty, err error)
	FileData(ctx context.Context, req *FileReq, opts ...http.CallOption) (rsp *FileDataRes, err error)
	FileDown(ctx context.Context, req *FileReq, opts ...http.CallOption) (rsp *FileDownRes, err error)
	ListFile(ctx context.Context, req *ListFileReq, opts ...http.CallOption) (rsp *ListFileReply, err error)
	ListVolume(ctx context.Context, req *ListVolumeReq, opts ...http.CallOption) (rsp *ListVolumeReply, err error)
	MoveAndCopyFile(ctx context.Context, req *MoveCopyFileReq, opts ...http.CallOption) (rsp *v1.Empty, err error)
	RenameFile(ctx context.Context, req *RenameFileReq, opts ...http.CallOption) (rsp *v1.Empty, err error)
}

type VolumeServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewVolumeServiceHTTPClient(client *http.Client) VolumeServiceHTTPClient {
	return &VolumeServiceHTTPClientImpl{client}
}

func (c *VolumeServiceHTTPClientImpl) CreateFile(ctx context.Context, in *CreateFileReq, opts ...http.CallOption) (*v1.Empty, error) {
	var out v1.Empty
	pattern := "/api/v1/volume/{id}/files"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.service.storage.v1.volume.VolumeService/CreateFile"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VolumeServiceHTTPClientImpl) CreateVolume(ctx context.Context, in *CreateVolumeReq, opts ...http.CallOption) (*Volume, error) {
	var out Volume
	pattern := "/api/v1/volume"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.service.storage.v1.volume.VolumeService/CreateVolume"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VolumeServiceHTTPClientImpl) DelFile(ctx context.Context, in *DelFileReq, opts ...http.CallOption) (*v1.Empty, error) {
	var out v1.Empty
	pattern := "/api/v1/volume/files/del"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.service.storage.v1.volume.VolumeService/DelFile"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VolumeServiceHTTPClientImpl) DeleteVolume(ctx context.Context, in *DeleteVolumeReq, opts ...http.CallOption) (*v1.Empty, error) {
	var out v1.Empty
	pattern := "/api/v1/volume/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.service.storage.v1.volume.VolumeService/DeleteVolume"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VolumeServiceHTTPClientImpl) FileData(ctx context.Context, in *FileReq, opts ...http.CallOption) (*FileDataRes, error) {
	var out FileDataRes
	pattern := "/api/v1/files/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.service.storage.v1.volume.VolumeService/FileData"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VolumeServiceHTTPClientImpl) FileDown(ctx context.Context, in *FileReq, opts ...http.CallOption) (*FileDownRes, error) {
	var out FileDownRes
	pattern := "/api/v1/volume/{id}/files/down"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.service.storage.v1.volume.VolumeService/FileDown"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VolumeServiceHTTPClientImpl) ListFile(ctx context.Context, in *ListFileReq, opts ...http.CallOption) (*ListFileReply, error) {
	var out ListFileReply
	pattern := "/api/v1/volume/{id}/files"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.service.storage.v1.volume.VolumeService/ListFile"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VolumeServiceHTTPClientImpl) ListVolume(ctx context.Context, in *ListVolumeReq, opts ...http.CallOption) (*ListVolumeReply, error) {
	var out ListVolumeReply
	pattern := "/api/v1/volume"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.service.storage.v1.volume.VolumeService/ListVolume"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VolumeServiceHTTPClientImpl) MoveAndCopyFile(ctx context.Context, in *MoveCopyFileReq, opts ...http.CallOption) (*v1.Empty, error) {
	var out v1.Empty
	pattern := "/api/v1/volume/files/copy-move"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.service.storage.v1.volume.VolumeService/MoveAndCopyFile"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *VolumeServiceHTTPClientImpl) RenameFile(ctx context.Context, in *RenameFileReq, opts ...http.CallOption) (*v1.Empty, error) {
	var out v1.Empty
	pattern := "/api/v1/volume/{id}/files/rename"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.service.storage.v1.volume.VolumeService/RenameFile"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
