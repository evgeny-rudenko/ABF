package api

import (
	"context"
	"log"

	pb "github.com/example/project/proto"
	"github.com/example/project/services"
)

// AuthHandler handles the authentication request.
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler creates a new AuthHandler instance.
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Authenticate handles the authentication request and returns the response.
func (h *AuthHandler) Authenticate(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	ok := h.authService.Authenticate(req.Login, req.Password, req.IP)

	return &pb.AuthResponse{
		Ok: ok,
	}, nil
}

// BlacklistHandler handles the IP blacklist operations.
type BlacklistHandler struct {
	blacklistService *services.BlacklistService
}

// NewBlacklistHandler creates a new BlacklistHandler instance.
func NewBlacklistHandler(blacklistService *services.BlacklistService) *BlacklistHandler {
	return &BlacklistHandler{
		blacklistService: blacklistService,
	}
}

// AddToBlacklist adds an IP subnet to the blacklist.
func (h *BlacklistHandler) AddToBlacklist(ctx context.Context, req *pb.BlacklistRequest) (*pb.BlacklistResponse, error) {
	err := h.blacklistService.AddToBlacklist(req.Subnet)
	if err != nil {
		log.Printf("Failed to add IP to blacklist: %v", err)
		return nil, err
	}

	return &pb.BlacklistResponse{}, nil
}

// RemoveFromBlacklist removes an IP subnet from the blacklist.
func (h *BlacklistHandler) RemoveFromBlacklist(ctx context.Context, req *pb.BlacklistRequest) (*pb.BlacklistResponse, error) {
	err := h.blacklistService.RemoveFromBlacklist(req.Subnet)
	if err != nil {
		log.Printf("Failed to remove IP from blacklist: %v", err)
		return nil, err
	}

	return &pb.BlacklistResponse{}, nil
}

// WhitelistHandler handles the IP whitelist operations.
type WhitelistHandler struct {
	whitelistService *services.WhitelistService
}

// NewWhitelistHandler creates a new WhitelistHandler instance.
func NewWhitelistHandler(whitelistService *services.WhitelistService) *WhitelistHandler {
	return &WhitelistHandler{
		whitelistService: whitelistService,
	}
}

// AddToWhitelist adds an IP subnet to the whitelist.
func (h *WhitelistHandler) AddToWhitelist(ctx context.Context, req *pb.WhitelistRequest) (*pb.WhitelistResponse, error) {
	err := h.whitelistService.AddToWhitelist(req.Subnet)
	if err != nil {
		log.Printf("Failed to add IP to whitelist: %v", err)
		return nil, err
	}

	return &pb.WhitelistResponse{}, nil
}

// RemoveFromWhitelist removes an IP subnet from the whitelist.
func (h *WhitelistHandler) RemoveFromWhitelist(ctx context.Context, req *pb.WhitelistRequest) (*pb.WhitelistResponse, error) {
	err := h.whitelistService.RemoveFromWhitelist(req.Subnet)
	if err != nil {
		log.Printf("Failed to remove IP from whitelist: %v", err)
		return nil, err
	}

	return &pb.WhitelistResponse{}, nil
}
