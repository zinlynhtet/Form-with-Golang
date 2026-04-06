package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
)

// Email configuration
const (
	recipientEmail = "mackk5504@gmail.com"
	smtpHost       = "smtp.gmail.com"
	smtpPort       = "587"
)

// Read sender credentials from environment variables
func getSenderEmail() string {
	email := os.Getenv("SENDER_EMAIL")
	if email == "" {
		email = "mackk5504@gmail.com" // Default: send from same email
	}
	return email
}

func getSenderAppPassword() string {
	return os.Getenv("GMAIL_APP_PASSWORD")
}

func sendEmail(firstName, lastName, email, phone, subject, message string) error {
	senderEmail := getSenderEmail()
	appPassword := getSenderAppPassword()

	if appPassword == "" {
		return fmt.Errorf("GMAIL_APP_PASSWORD environment variable is not set")
	}

	// Build a nicely formatted email body
	subjectMap := map[string]string{
		"general":     "General Inquiry",
		"support":     "Technical Support",
		"feedback":    "Feedback",
		"partnership": "Partnership",
		"other":       "Other",
	}

	subjectLabel := subjectMap[subject]
	if subjectLabel == "" {
		subjectLabel = subject
	}

	// Compose the email with HTML formatting
	emailSubject := fmt.Sprintf("📬 New Contact Form: %s from %s %s", subjectLabel, firstName, lastName)

	htmlBody := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <style>
        body { font-family: 'Segoe UI', Arial, sans-serif; background: #f0f2f5; margin: 0; padding: 20px; }
        .container { max-width: 600px; margin: 0 auto; background: #ffffff; border-radius: 16px; overflow: hidden; box-shadow: 0 4px 24px rgba(0,0,0,0.08); }
        .header { background: linear-gradient(135deg, #6C63FF, #4F46E5); padding: 32px; text-align: center; }
        .header h1 { color: #ffffff; font-size: 22px; margin: 0; font-weight: 600; }
        .header p { color: rgba(255,255,255,0.8); font-size: 14px; margin-top: 8px; }
        .body { padding: 32px; }
        .field { margin-bottom: 20px; }
        .field-label { font-size: 11px; text-transform: uppercase; letter-spacing: 1px; color: #6B7280; font-weight: 600; margin-bottom: 6px; }
        .field-value { font-size: 15px; color: #1F2937; padding: 12px 16px; background: #F9FAFB; border-radius: 10px; border-left: 3px solid #6C63FF; }
        .message-box { font-size: 15px; color: #1F2937; padding: 16px; background: #F9FAFB; border-radius: 10px; border-left: 3px solid #06D6A0; line-height: 1.6; white-space: pre-wrap; }
        .footer { text-align: center; padding: 20px 32px; background: #F9FAFB; border-top: 1px solid #E5E7EB; }
        .footer p { font-size: 12px; color: #9CA3AF; margin: 0; }
        .badge { display: inline-block; padding: 4px 12px; background: rgba(108,99,255,0.1); color: #6C63FF; border-radius: 20px; font-size: 12px; font-weight: 600; }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>📬 New Contact Form Submission</h1>
            <p>You received a new message from your website</p>
        </div>
        <div class="body">
            <div class="field">
                <div class="field-label">Full Name</div>
                <div class="field-value">%s %s</div>
            </div>
            <div class="field">
                <div class="field-label">Email Address</div>
                <div class="field-value"><a href="mailto:%s" style="color: #6C63FF; text-decoration: none;">%s</a></div>
            </div>
            <div class="field">
                <div class="field-label">Phone Number</div>
                <div class="field-value">%s</div>
            </div>
            <div class="field">
                <div class="field-label">Subject</div>
                <div class="field-value"><span class="badge">%s</span></div>
            </div>
            <div class="field">
                <div class="field-label">Message</div>
                <div class="message-box">%s</div>
            </div>
        </div>
        <div class="footer">
            <p>This email was sent from your Contact Form application</p>
        </div>
    </div>
</body>
</html>`,
		firstName, lastName,
		email, email,
		phone,
		subjectLabel,
		message,
	)

	// Build MIME email
	var emailBuilder strings.Builder
	emailBuilder.WriteString(fmt.Sprintf("From: %s\r\n", senderEmail))
	emailBuilder.WriteString(fmt.Sprintf("To: %s\r\n", recipientEmail))
	emailBuilder.WriteString(fmt.Sprintf("Subject: %s\r\n", emailSubject))
	emailBuilder.WriteString("MIME-Version: 1.0\r\n")
	emailBuilder.WriteString("Content-Type: text/html; charset=\"UTF-8\"\r\n")
	emailBuilder.WriteString("\r\n")
	emailBuilder.WriteString(htmlBody)

	// Authenticate and send
	auth := smtp.PlainAuth("", senderEmail, appPassword, smtpHost)
	addr := smtpHost + ":" + smtpPort

	err := smtp.SendMail(addr, auth, senderEmail, []string{recipientEmail}, []byte(emailBuilder.String()))
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	// ParseMultipartForm handles multipart/form-data (from fetch + FormData)
	// The 10 << 20 = 10MB max memory for form data
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		// Fallback to ParseForm for url-encoded submissions
		if err2 := r.ParseForm(); err2 != nil {
			http.Error(w, fmt.Sprintf("ParseForm() err: %v", err2), http.StatusBadRequest)
			return
		}
	}

	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	subject := r.FormValue("subject")
	message := r.FormValue("message")

	fmt.Printf("=== New Contact Form Submission ===\n")
	fmt.Printf("Name:    %s %s\n", firstName, lastName)
	fmt.Printf("Email:   %s\n", email)
	fmt.Printf("Phone:   %s\n", phone)
	fmt.Printf("Subject: %s\n", subject)
	fmt.Printf("Message: %s\n", message)

	// Send email notification
	if err := sendEmail(firstName, lastName, email, phone, subject, message); err != nil {
		fmt.Printf("Email error: %v\n", err)
		fmt.Printf("===================================\n\n")
		// Still return success to the user (form data is logged)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Thank you, %s! Your message has been received email notification failed).", firstName)
		return
	}

	fmt.Printf("Email sent to %s\n", recipientEmail)
	fmt.Printf("===================================\n\n")

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Thank you, %s! Your message has been received and emailed.", firstName)
}

func main() {
	// Check for app password
	if getSenderAppPassword() == "" {
		fmt.Println("WARNING: GMAIL_APP_PASSWORD environment variable is not set!")
		fmt.Println("Email notifications will NOT work until you set it.")
		fmt.Println("Run: $env:GMAIL_APP_PASSWORD=\"gjdg pxxp hlre xrfa\"")
		fmt.Println()
	} else {
		fmt.Printf("Email configured: notifications will be sent to %s\n", recipientEmail)
	}

	fileServer := http.FileServer(http.Dir("./static"))	
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("🚀 Server starting at http://localhost:8080\n")
	fmt.Printf("📋 Contact form at http://localhost:8080/form.html\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
