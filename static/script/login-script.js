// Current section
        let currentSection = 0;
        let currentStep = 1;
        let currentSocialProvider = '';
        
        // Switch between sections
        function switchSection(index) {
            const navLinks = document.querySelectorAll('.nav-link');
            navLinks.forEach((link, i) => {
                if (i === index) {
                    link.classList.add('active');
                } else {
                    link.classList.remove('active');
                }
            });
            
            document.getElementById('authWrapper').style.transform = `translateX(-${index * 20}%)`;
            currentSection = index;
        }
        
        // Toggle password visibility
        function togglePassword(inputId, icon) {
            const input = document.getElementById(inputId);
            const type = input.getAttribute('type') === 'password' ? 'text' : 'password';
            input.setAttribute('type', type);
            
            if (type === 'text') {
                icon.innerHTML = '<i class="fas fa-eye-slash"></i>';
            } else {
                icon.innerHTML = '<i class="fas fa-eye"></i>';
            }
        }
        
        // Toggle checkbox
        function toggleCheckbox(checkbox) {
            checkbox.classList.toggle('checked');
            const input = checkbox.querySelector('input');
            input.checked = !input.checked;
        }
        
        // Toggle theme
        function toggleTheme(toggle) {
            toggle.classList.toggle('dark');
            // Additional theme switching logic would go here
        }
        
        // Next step in signup
        function nextStep(step) {
            // Validate current step
            if (step === 1) {
                if (!validateStep1()) return;
            } else if (step === 2) {
                if (!validateStep2()) return;
            } else if (step === 3) {
                if (!validateStep3()) return;
            }
            
            // Hide current step
            document.getElementById(`step${step}`).classList.remove('active');
            
            // Mark current step as completed
            document.getElementById(`step${step}Circle`).classList.add('completed');
            
            // Show next step
            document.getElementById(`step${step + 1}`).classList.add('active');
            
            // Mark next step as active
            document.getElementById(`step${step + 1}Circle`).classList.add('active');
            
            currentStep = step + 1;
        }
        
        // Previous step in signup
        function prevStep(step) {
            // Hide current step
            document.getElementById(`step${step}`).classList.remove('active');
            
            // Remove active class from current step
            document.getElementById(`step${step}Circle`).classList.remove('active');
            
            // Show previous step
            document.getElementById(`step${step - 1}`).classList.add('active');
            
            // Remove completed class from previous step
            document.getElementById(`step${step - 1}Circle`).classList.remove('completed');
            
            currentStep = step - 1;
        }
        
        // Validate step 1
        function validateStep1() {
            let isValid = true;
            
            // Validate email
            const email = document.getElementById('signupEmail').value;
            const emailGroup = document.getElementById('signupEmailGroup');
            
            if (!validateEmail(email)) {
                emailGroup.classList.add('error');
                isValid = false;
            } else {
                emailGroup.classList.remove('error');
            }
            
            // Validate username
            const username = document.getElementById('signupUsername').value;
            const usernameGroup = document.getElementById('signupUsernameGroup');
            
            if (username.length < 4) {
                usernameGroup.classList.add('error');
                isValid = false;
            } else {
                usernameGroup.classList.remove('error');
            }
            
            // Validate password
            const password = document.getElementById('signupPassword').value;
            const passwordGroup = document.getElementById('signupPasswordGroup');
            
            if (password.length < 8) {
                passwordGroup.classList.add('error');
                isValid = false;
            } else {
                passwordGroup.classList.remove('error');
            }
            
            // Validate confirm password
            const confirmPassword = document.getElementById('confirmPassword').value;
            const confirmPasswordGroup = document.getElementById('confirmPasswordGroup');
            
            if (password !== confirmPassword) {
                confirmPasswordGroup.classList.add('error');
                isValid = false;
            } else {
                confirmPasswordGroup.classList.remove('error');
            }
            
            return isValid;
        }
        
        // Validate step 2
        function validateStep2() {
            let isValid = true;
            
            // Validate first name
            const firstName = document.getElementById('firstName').value;
            const firstNameGroup = document.getElementById('firstNameGroup');
            
            if (firstName.trim() === '') {
                firstNameGroup.classList.add('error');
                isValid = false;
            } else {
                firstNameGroup.classList.remove('error');
            }
            
            // Validate last name
            const lastName = document.getElementById('lastName').value;
            const lastNameGroup = document.getElementById('lastNameGroup');
            
            if (lastName.trim() === '') {
                lastNameGroup.classList.add('error');
                isValid = false;
            } else {
                lastNameGroup.classList.remove('error');
            }
            
            return isValid;
        }
        
        // Validate step 3
        function validateStep3() {
            let isValid = true;
            
            // Validate security questions
            const question1 = document.getElementById('securityQuestion1').value;
            const answer1 = document.getElementById('securityAnswer1').value;
            const question2 = document.getElementById('securityQuestion2').value;
            const answer2 = document.getElementById('securityAnswer2').value;
            
            if (question1 === '' || answer1.trim() === '' || question2 === '' || answer2.trim() === '') {
                isValid = false;
                showNotification('error', 'Error', 'Please fill in all security questions and answers.');
            }
            
            return isValid;
        }
        
        // Handle login
        function handleLogin() {
            const email = document.getElementById('loginEmail').value;
            const password = document.getElementById('loginPassword').value;
            const emailGroup = document.getElementById('loginEmailGroup');
            const passwordGroup = document.getElementById('loginPasswordGroup');
            let isValid = true;
            
            // Fixed: Improved email validation
            if (!validateEmail(email)) {
                emailGroup.classList.add('error');
                isValid = false;
            } else {
                emailGroup.classList.remove('error');
            }
            
            if (password.trim() === '') {
                passwordGroup.classList.add('error');
                isValid = false;
            } else {
                passwordGroup.classList.remove('error');
            }
            
            if (isValid) {
                showNotification('success', 'Success', 'You have successfully logged in!');
            }
        }
        
        // Handle signup
        function handleSignup() {
            const termsCheckbox = document.getElementById('termsCheckbox');
        
            if (!termsCheckbox.classList.contains('checked')) {
                showNotification('error', 'Error', 'You must agree to the Terms of Service and Privacy Policy.');
                return;
            }

            const form = document.getElementById('signupForm');
            const formData = new FormData(form); 

            fetch('/register', { 
                method: 'POST',
                body: formData 
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('خطا در ثبت‌نام');
                }
                return response.text();
            })
            .then(data => {
                showNotification('success', 'Success', 'Your account has been created successfully!');
                setTimeout(() => {
                    switchSection(0); 
                }, 2000);
            })
            .catch(error => {
                showNotification('error', 'Error', `خطا: ${error.message}`);
            });
        }

document.getElementById("signupBtn").removeEventListener("click", function() {
    location.reload();
});
document.getElementById("signupBtn").addEventListener("click", handleSignup);
// setTimeout(() => {
//                 switchSection(0);
//             }, 2000);
// document.getElementById("signupBtn").addEventListener("click", handleSignup); 
        
        // Show notification
        function showNotification(type, title, message) {
            const notification = document.getElementById('notification');
            const notificationTitle = notification.querySelector('.notification-title');
            const notificationMessage = notification.querySelector('.notification-message');
            const notificationIcon = notification.querySelector('.notification-icon i');
            
            notification.className = 'notification';
            notification.classList.add(type);
            
            notificationTitle.textContent = title;
            notificationMessage.textContent = message;
            
            if (type === 'success') {
                notificationIcon.className = 'fas fa-check-circle';
            } else if (type === 'error') {
                notificationIcon.className = 'fas fa-exclamation-circle';
            } else if (type === 'info') {
                notificationIcon.className = 'fas fa-info-circle';
            } else if (type === 'warning') {
                notificationIcon.className = 'fas fa-exclamation-triangle';
            }
            
            notification.classList.add('show');
            
            setTimeout(() => {
                closeNotification();
            }, 5000);
        }
        
        // Close notification
        function closeNotification() {
            const notification = document.getElementById('notification');
            notification.classList.remove('show');
        }
        
        // Select account type
        function selectAccountType(element, type) {
            const personalAccount = document.getElementById('personalAccount');
            const businessAccount = document.getElementById('businessAccount');
            const personalRadio = document.getElementById('personalRadio');
            const businessRadio = document.getElementById('businessRadio');
            
            if (type === 'personal') {
                personalAccount.style.background = 'rgba(79, 70, 229, 0.1)';
                businessAccount.style.background = 'rgba(255, 255, 255, 0.05)';
                personalRadio.querySelector('div').style.display = 'block';
                businessRadio.querySelector('div').style.display = 'none';
            } else {
                personalAccount.style.background = 'rgba(255, 255, 255, 0.05)';
                businessAccount.style.background = 'rgba(79, 70, 229, 0.1)';
                personalRadio.querySelector('div').style.display = 'none';
                businessRadio.querySelector('div').style.display = 'block';
            }
        }
        
        // Password strength checker
        document.getElementById('signupPassword').addEventListener('input', function() {
            const password = this.value;
            const strengthMeter = document.getElementById('strengthMeter');
            const strengthText = document.getElementById('strengthText');
            const strengthScore = document.getElementById('strengthScore');
            
            // Check requirements
            const hasLength = password.length >= 8;
            const hasUppercase = /[A-Z]/.test(password);
            const hasLowercase = /[a-z]/.test(password);
            const hasNumber = /[0-9]/.test(password);
            const hasSpecial = /[^A-Za-z0-9]/.test(password);
            
            // Update requirement indicators
            updateRequirement('req-length', hasLength);
            updateRequirement('req-uppercase', hasUppercase);
            updateRequirement('req-lowercase', hasLowercase);
            updateRequirement('req-number', hasNumber);
            updateRequirement('req-special', hasSpecial);
            
            // Calculate score
            let score = 0;
            if (hasLength) score += 20;
            if (hasUppercase) score += 20;
            if (hasLowercase) score += 20;
            if (hasNumber) score += 20;
            if (hasSpecial) score += 20;
            
            // Update meter
            strengthMeter.style.width = `${score}%`;
            strengthScore.textContent = `${score}/100`;
            
            // Update class and text
            strengthMeter.className = 'strength-meter-fill';
            if (score <= 20) {
                strengthMeter.classList.add('weak');
                strengthText.textContent = 'Very Weak';
            } else if (score <= 40) {
                strengthMeter.classList.add('weak');
                strengthText.textContent = 'Weak';
            } else if (score <= 60) {
                strengthMeter.classList.add('medium');
                strengthText.textContent = 'Medium';
            } else if (score <= 80) {
                strengthMeter.classList.add('good');
                strengthText.textContent = 'Good';
            } else {
                strengthMeter.classList.add('strong');
                strengthText.textContent = 'Strong';
            }
        });
        
        // Update password requirement indicator
        function updateRequirement(id, isValid) {
            const req = document.getElementById(id);
            const icon = req.querySelector('.requirement-icon');
            
            if (isValid) {
                icon.className = 'requirement-icon success';
                icon.innerHTML = '<i class="fas fa-check"></i>';
            } else {
                icon.className = 'requirement-icon pending';
                icon.innerHTML = '<i class="fas fa-circle"></i>';
            }
        }
        
        // Fixed: Improved email validation function
        function validateEmail(email) {
            const re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
            return re.test(String(email).toLowerCase());
        }
        
        // Fixed: Password confirmation validation
        document.getElementById('confirmPassword').addEventListener('input', function() {
            const password = document.getElementById('signupPassword').value;
            const confirmPassword = this.value;
            const confirmPasswordGroup = document.getElementById('confirmPasswordGroup');
            
            if (password !== confirmPassword) {
                confirmPasswordGroup.classList.add('error');
            } else {
                confirmPasswordGroup.classList.remove('error');
            }
        });
        
        // Show forgot password modal
        function showForgotPasswordModal() {
            document.getElementById('forgotPasswordModal').classList.add('show');
        }
        
        // Close modal
        function closeModal(modalId) {
            document.getElementById(modalId).classList.remove('show');
        }
        
        // Send password reset
        function sendPasswordReset() {
            const email = document.getElementById('resetEmail').value;
            const phone = document.getElementById('resetPhone').value;
            const emailGroup = document.getElementById('resetEmailGroup');
            const phoneGroup = document.getElementById('resetPhoneGroup');
            let isValid = true;
            
            if (!validateEmail(email)) {
                emailGroup.classList.add('error');
                isValid = false;
            } else {
                emailGroup.classList.remove('error');
            }
            
            if (phone.trim() === '') {
                phoneGroup.classList.add('error');
                isValid = false;
            } else {
                phoneGroup.classList.remove('error');
            }
            
            if (isValid) {
                closeModal('forgotPasswordModal');
                showNotification('success', 'Email Sent', 'Password reset instructions have been sent to your email.');
            }
        }
        
        // Show social login modal
        function showSocialLoginModal(provider) {
            currentSocialProvider = provider;
            const modal = document.getElementById('socialLoginModal');
            const title = document.getElementById('socialModalTitle');
            const text = document.getElementById('socialModalText');
            
            // Hide all account lists
            document.getElementById('googleAccounts').style.display = 'none';
            document.getElementById('facebookAccounts').style.display = 'none';
            document.getElementById('twitterAccounts').style.display = 'none';
            document.getElementById('appleAccounts').style.display = 'none';
            
            // Show the appropriate account list
            document.getElementById(`${provider}Accounts`).style.display = 'flex';
            
            // Update title and text
            if (provider === 'google') {
                title.textContent = 'Choose a Google Account';
                text.textContent = 'Select an account to continue with Google';
            } else if (provider === 'facebook') {
                title.textContent = 'Choose a Facebook Account';
                text.textContent = 'Select an account to continue with Facebook';
            } else if (provider === 'twitter') {
                title.textContent = 'Choose a Twitter Account';
                text.textContent = 'Select an account to continue with Twitter';
            } else if (provider === 'apple') {
                title.textContent = 'Choose an Apple Account';
                text.textContent = 'Select an account to continue with Apple';
            }
            
            modal.classList.add('show');
        }
        
        // Login with social account
        function loginWithSocialAccount(provider, account) {
            closeModal('socialLoginModal');
            showNotification('success', 'Success', `You have successfully logged in with ${provider} as ${account}`);
        }
        
        // Initialize event listeners
        document.addEventListener('DOMContentLoaded', function() {
            // Email validation on input
            document.getElementById('loginEmail').addEventListener('input', function() {
                const emailGroup = document.getElementById('loginEmailGroup');
                if (validateEmail(this.value)) {
                    emailGroup.classList.remove('error');
                }
            });
            
            document.getElementById('signupEmail').addEventListener('input', function() {
                const emailGroup = document.getElementById('signupEmailGroup');
                if (validateEmail(this.value)) {
                    emailGroup.classList.remove('error');
                }
            });
            
            // Use another account link
            document.getElementById('useAnotherAccount').addEventListener('click', function(e) {
                e.preventDefault();
                closeModal('socialLoginModal');
            });
        });