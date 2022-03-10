package main

import (
	"github.com/OpenCal-FYDP/CalendarEventManagement/internal/service"
	"github.com/OpenCal-FYDP/CalendarEventManagement/rpc"
	"log"
	"net/http"
)

func main() {
	svc := service.New()
	server := rpc.NewCalendarEventManagementServiceServer(svc)
	log.Fatal(http.ListenAndServe(":8080", server))
}
