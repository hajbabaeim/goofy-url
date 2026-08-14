package http

func CreateShortURL(c echo.Context, svc *service.ShortenerService) error {
	var req struct{ URL string }
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, "invalid request")
	}

	shortURL, err := svc.Create(req.URL)
	if err != nil {
		return c.JSON(500, "server error")
	}
	return c.JSON(201, shortURL)
}
