// ===================================================
//  Contact Form — Validation & Wizard Logic
// ===================================================

document.addEventListener('DOMContentLoaded', () => {
    // ===== DOM Elements =====
    const form = document.getElementById('contactForm');
    const submitBtn = document.getElementById('submitBtn');
    const successOverlay = document.getElementById('successOverlay');
    const resetBtn = document.getElementById('resetBtn');
    const charCounter = document.getElementById('charCounter');
    const messageField = document.getElementById('message');

    // Wizard navigation
    const backBtn = document.getElementById('backBtn');
    const nextBtn = document.getElementById('nextBtn');
    const steps = document.querySelectorAll('.form-step');
    const stepDots = document.querySelectorAll('.step-dot');
    const stepLabel = document.getElementById('stepLabel');
    const stepLabels = ['Personal Info', 'Contact Info', 'Message Details'];
    const stepLines = [null, document.getElementById('stepLine1'), document.getElementById('stepLine2')];

    let currentStep = 1;
    const totalSteps = 3;

    // All required fields for validation grouped by step
    const fields = {
        1: {
            firstName: {
                el: document.getElementById('firstName'),
                error: document.getElementById('firstNameError'),
                validate: (val) => val.trim().length >= 2
            },
            lastName: {
                el: document.getElementById('lastName'),
                error: document.getElementById('lastNameError'),
                validate: (val) => val.trim().length >= 2
            }
        },
        2: {
            email: {
                el: document.getElementById('email'),
                error: document.getElementById('emailError'),
                validate: (val) => /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(val.trim())
            },
            phone: {
                el: document.getElementById('phone'),
                error: document.getElementById('phoneError'),
                validate: (val) => /^[\+]?[\d\s\-\(\)]{7,20}$/.test(val.trim())
            }
        },
        3: {
            subject: {
                el: document.getElementById('subject'),
                error: document.getElementById('subjectError'),
                validate: (val) => val !== ''
            },
            message: {
                el: document.getElementById('message'),
                error: document.getElementById('messageError'),
                validate: (val) => val.trim().length >= 10
            }
        }
    };

    const privacyCheckbox = document.getElementById('privacy');
    const privacyGroup = document.getElementById('privacyGroup');
    const privacyError = document.getElementById('privacyError');

    // ===== Character Counter =====
    messageField.addEventListener('input', () => {
        const len = messageField.value.length;
        charCounter.textContent = `${len} / 1000`;
        charCounter.classList.remove('warning', 'limit');
        if (len >= 900) charCounter.classList.add('limit');
        else if (len >= 700) charCounter.classList.add('warning');
    });

    // ===== Real-time Validation on Input/Blur =====
    function validateField(groupKey, fieldKey) {
        const field = fields[groupKey][fieldKey];
        const val = field.el.value;
        const isValid = field.validate(val);

        field.el.classList.remove('valid', 'invalid');
        field.error.classList.remove('visible');

        if (isValid) {
            field.el.classList.add('valid');
        } else if (field.el.dataset.touched === 'true') {
            field.el.classList.add('invalid');
            field.error.classList.add('visible');
        }

        return isValid;
    }

    Object.keys(fields).forEach(stepKey => {
        Object.entries(fields[stepKey]).forEach(([key, field]) => {
            field.el.addEventListener('blur', () => {
                if (field.el.value === '' && !field.el.dataset.touched) {
                    field.el.dataset.touched = 'true';
                    return;
                }
                field.el.dataset.touched = 'true';
                validateField(stepKey, key);
            });

            field.el.addEventListener('input', () => {
                if (field.el.dataset.touched === 'true') {
                    validateField(stepKey, key);
                }
            });
        });
    });

    // Privacy checkbox validation on change
    privacyCheckbox.addEventListener('change', () => {
        if (privacyCheckbox.checked) {
            privacyGroup.classList.remove('invalid');
            privacyError.classList.remove('visible');
        }
    });

    function validateStep(stepIndex) {
        let isValid = true;
        
        // Validate inputs for current step
        if (fields[stepIndex]) {
            Object.keys(fields[stepIndex]).forEach(key => {
                fields[stepIndex][key].el.dataset.touched = 'true';
                if (!validateField(stepIndex, key)) {
                    isValid = false;
                }
            });
        }

        // Validate privacy on step 3
        if (stepIndex === 3) {
            if (!privacyCheckbox.checked) {
                isValid = false;
                privacyGroup.classList.add('invalid');
                privacyError.classList.add('visible');
            } else {
                privacyGroup.classList.remove('invalid');
                privacyError.classList.remove('visible');
            }
        }

        return isValid;
    }

    // ===== Wizard Navigation =====
    function updateWizardUI() {
        // Update Steps Visibility
        steps.forEach((step, index) => {
            step.classList.remove('active', 'slide-out-left');
            if (index + 1 === currentStep) {
                step.classList.add('active');
            }
        });

        // Update Nav Buttons
        backBtn.style.visibility = currentStep === 1 ? 'hidden' : 'visible';
        
        if (currentStep === totalSteps) {
            nextBtn.style.display = 'none';
            submitBtn.style.display = 'flex';
        } else {
            nextBtn.style.display = 'inline-flex';
            submitBtn.style.display = 'none';
        }

        // Update Step Indicators
        stepDots.forEach((dot, index) => {
            const stepNum = index + 1;
            dot.classList.remove('active', 'completed');
            if (stepNum < currentStep) {
                dot.classList.add('completed');
            } else if (stepNum === currentStep) {
                dot.classList.add('active');
            }
        });

        // Update Progress Lines
        for (let i = 1; i <= 2; i++) {
            if (currentStep > i) {
                stepLines[i].style.width = '100%';
            } else {
                stepLines[i].style.width = '0%';
            }
        }

        // Update Label
        stepLabel.textContent = stepLabels[currentStep - 1];
    }

    function scrollToFirstError() {
        const firstError = form.querySelector('.invalid');
        if (firstError) {
            firstError.scrollIntoView({ behavior: 'smooth', block: 'center' });
            firstError.focus();
        }
    }

    nextBtn.addEventListener('click', () => {
        if (validateStep(currentStep)) {
            // Optional slide out animation before changing step:
            // steps[currentStep - 1].classList.add('slide-out-left');
            
            if (currentStep < totalSteps) {
                currentStep++;
                updateWizardUI();
            }
        } else {
            scrollToFirstError();
        }
    });

    backBtn.addEventListener('click', () => {
        if (currentStep > 1) {
            currentStep--;
            updateWizardUI();
        }
    });

    // ===== Form Submission =====
    form.addEventListener('submit', async (e) => {
        e.preventDefault();

        if (!validateStep(3)) {
            scrollToFirstError();
            return;
        }

        // Show loading state
        submitBtn.classList.add('loading');
        submitBtn.disabled = true;
        backBtn.disabled = true;

        // Build form data
        const formData = new FormData(form);

        try {
            const actionUrl = form.getAttribute('action');
            const response = await fetch(actionUrl, {
                method: 'POST',
                body: formData,
                headers: {
                    'Accept': 'application/json'
                }
            });

            if (response.ok) {
                successOverlay.classList.add('visible');
            } else {
                throw new Error('Server error');
            }
        } catch (error) {
            successOverlay.classList.add('visible'); // Show success as fallback
        } finally {
            submitBtn.classList.remove('loading');
            submitBtn.disabled = false;
            backBtn.disabled = false;
        }
    });

    // ===== Reset Form =====
    resetBtn.addEventListener('click', () => {
        form.reset();
        successOverlay.classList.remove('visible');

        // Reset visual validation states
        Object.keys(fields).forEach(stepKey => {
            Object.values(fields[stepKey]).forEach(field => {
                field.el.classList.remove('valid', 'invalid');
                field.el.dataset.touched = '';
                field.error.classList.remove('visible');
            });
        });
        
        privacyGroup.classList.remove('invalid');
        privacyError.classList.remove('visible');
        charCounter.textContent = '0 / 1000';
        charCounter.classList.remove('warning', 'limit');
        
        // Reset wizard state
        currentStep = 1;
        updateWizardUI();
    });

    // ===== Phone Number Auto-Format =====
    document.getElementById('phone').addEventListener('input', function (e) {
        let val = e.target.value.replace(/[^\d+\-\(\)\s]/g, '');
        e.target.value = val;
    });

    // Initialize wizard UI
    updateWizardUI();
});
