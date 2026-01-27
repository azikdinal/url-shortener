package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorten/internal/domain"
	"shorten/internal/usecase"
)

type Handler struct {
	usecase    usecase.Usecase
	publicHost string
}

func NewHandler(uc usecase.Usecase, ph string) *Handler {
	return &Handler{
		usecase:    uc,
		publicHost: ph,
	}
}

func (h *Handler) GetByCode(c *gin.Context) {
	ctx := c.Request.Context()

	sc, err := domain.NewShortCode(c.Param("code"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "short code is invalid"})
	}

	fu, err := h.usecase.GetByCode(ctx, sc)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "link not found"})
		return
	}
	link := domain.NewLink(sc, fu)

	c.JSON(http.StatusOK, toGetResponse(link, h.publicHost))
}

func (h *Handler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	fu, err := domain.NewFullURL(req.FullURL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid fullURL"})
		return
	}

	sc, err := h.usecase.Create(ctx, fu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create link"})
		return
	}

	c.JSON(http.StatusOK, toCreateResponse(sc, h.publicHost))
}
