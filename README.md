# 📬 Professional Contact Form

A beautifully crafted, modern contact form built with **Go** and vanilla **HTML/CSS/JavaScript**. Features a stunning glassmorphism dark theme, real-time form validation, email notifications, and full GitHub Pages deployment support.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go&logoColor=white)
![HTML5](https://img.shields.io/badge/HTML5-E34F26?style=flat-square&logo=html5&logoColor=white)
![CSS3](https://img.shields.io/badge/CSS3-1572B6?style=flat-square&logo=css3&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=flat-square&logo=javascript&logoColor=black)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)

---

## ✨ Features

| Feature | Description |
|---------|-------------|
| 🎨 **Glassmorphism Design** | Premium dark theme with frosted-glass cards, floating orbs, and gradient accents |
| ✅ **Real-time Validation** | Instant field validation on blur with green/red visual feedback |
| 📊 **Progress Tracker** | Live progress bar showing completed fields (0/6 → 6/6) |
| 📧 **Email Notifications** | Form submissions sent directly to your email via Gmail SMTP |
| 🔒 **Privacy Consent** | Mandatory privacy policy checkbox before submission |
| 📱 **Fully Responsive** | Looks great on desktop, tablet, and mobile devices |
| ⚡ **Micro-animations** | Smooth transitions, hover effects, and animated success overlay |
| 🔤 **Character Counter** | Message field shows live character count with color warnings |
| 📞 **Phone Auto-format** | Automatically strips invalid characters from phone input |
| 🌐 **GitHub Pages Ready** | Static version deployable to GitHub Pages with Formspree |

---

## 📁 Project Structure

```
Form-with-Golang/
├── main.go                  # Go backend server with email support
├── README.md                # This documentation
├── static/                  # Static frontend files (also served by GitHub Pages)
│   ├── index.html           # Landing page with hero section
│   ├── form.html            # Professional contact form page
│   ├── css/
│   │   ├── shared.css       # Design system: variables, reset, base styles
│   │   ├── index.css        # Landing page specific styles
│   │   └── form.css         # Contact form specific styles
│   └── js/
│       └── form.js          # Form validation & interaction logic
```

---

## 🚀 Getting Started

### Prerequisites

- [Go 1.21+](https://go.dev/dl/) installed
- A Gmail account (for email notifications)

### 1. Clone the Repository

```bash
git clone https://github.com/zinlynhtet/Form-with-Golang.git
cd Form-with-Golang
```

### 2. Run the Server (without email)

```bash
go run main.go
```

The server starts at **http://localhost:8080**

- Landing page: http://localhost:8080/
- Contact form: http://localhost:8080/form.html

### 3. Run with Email Notifications

To receive form submissions in your email, set up a Gmail App Password:

1. Go to [Google Account Security](https://myaccount.google.com/security)
2. Enable **2-Step Verification** if not already enabled
3. Search for **App Passwords** → Create one for "Mail"
4. Copy the 16-character password

Then run:

**Windows (PowerShell):**
```powershell
$env:GMAIL_APP_PASSWORD="your-16-char-app-password"
$env:SENDER_EMAIL="your-email@gmail.com"
go run main.go
```

**macOS / Linux:**
```bash
export GMAIL_APP_PASSWORD="your-16-char-app-password"
export SENDER_EMAIL="your-email@gmail.com"
go run main.go
```

---

## 🌐 Deploy to GitHub Pages

GitHub Pages serves static files only (no Go backend), so the form uses [Formspree](https://formspree.io) to handle submissions and send them to your email.

### Step 1: Set Up Formspree

1. Go to [formspree.io](https://formspree.io) and sign up (free tier: 50 submissions/month)
2. Create a new form → Enter your email `example@gmail.com`
3. Copy your form endpoint URL (e.g., `https://formspree.io/f/xAbCdEfG`)

### Step 2: Update the Form Action

In `static/form.html`, update the form action with your Formspree endpoint:

```html
<form action="https://formspree.io/f/YOUR_FORM_ID" method="post" id="contactForm" novalidate>
```

### Step 3: Configure GitHub Pages

1. Go to your GitHub repository: **Settings** → **Pages**
2. Under **Build and deployment** → **Source**, select: **GitHub Actions**
3. Push the code to GitHub. The `.github/workflows/deploy.yml` file included in this repository will automatically build and deploy your `static` folder!

### Step 4: Push Your Changes

```bash
git add .
git commit -m "Deploy contact form with GitHub actions"
git push origin main
```

Your site will be live at: **https://zinlynhtet.github.io/Form-with-Golang/**

> **Note:** It may take 1-2 minutes for GitHub Pages to build and deploy.

---

## 🎨 Design System

The CSS architecture uses a modular approach with CSS custom properties:

### Color Palette

| Variable | Value | Usage |
|----------|-------|-------|
| `--primary` | `#6C63FF` | Buttons, focus states, accents |
| `--primary-dark` | `#4F46E5` | Gradients, hover states |
| `--accent` | `#06D6A0` | Success states, checkmarks |
| `--bg-dark` | `#0F0E17` | Page background |
| `--error` | `#FF6B6B` | Validation errors |
| `--text-primary` | `#FFFFFE` | Main text |
| `--text-muted` | `#6B6D7B` | Secondary/helper text |

### CSS Files

| File | Purpose |
|------|---------|
| `shared.css` | CSS reset, variables, base styles, animated background, shared keyframes |
| `index.css` | Hero section, floating icon, CTA button |
| `form.css` | Form fields, validation states, submit button, success overlay, responsive |

---

## 🔧 Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `GMAIL_APP_PASSWORD` | _(none)_ | Gmail App Password for SMTP authentication |
| `SENDER_EMAIL` | `example@gmail.com` | Email address used to send notifications |

### Customization

- **Recipient email**: Change `recipientEmail` in `main.go`
- **Form fields**: Edit the HTML in `form.html` and update validation in `js/form.js`
- **Colors**: Modify CSS variables in `css/shared.css` under `:root`
- **Formspree endpoint**: Update the `action` attribute in `form.html`

---

## 📋 Form Fields

All fields are **mandatory** and validated in real-time:

| Field | Validation Rule |
|-------|----------------|
| First Name | Min 2 characters |
| Last Name | Min 2 characters |
| Email | Valid email format (regex) |
| Phone | 7-20 digits, allows +, -, (), spaces |
| Subject | Must select an option |
| Message | Min 10 characters, max 1000 |
| Privacy | Must be checked |

---

## 🛠️ Tech Stack

- **Backend**: Go `net/http` standard library
- **Email**: Go `net/smtp` with Gmail SMTP
- **Frontend**: Vanilla HTML5, CSS3, JavaScript (ES6+)
- **Fonts**: [Inter](https://fonts.google.com/specimen/Inter) via Google Fonts
- **Static Hosting**: GitHub Pages
- **Form Service**: [Formspree](https://formspree.io) (for GitHub Pages deployment)

---

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/zinlynhtet">Zin Linn Htet</a>
</p>
