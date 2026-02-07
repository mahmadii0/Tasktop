

// ============================================
// 1. SMOOTH SCROLL FOR NAVIGATION LINKS
// ============================================
// When user clicks on menu links, page scrolls smoothly to sections
document.addEventListener('DOMContentLoaded', function() {
    
    // Get all menu links
    const menuLinks = document.querySelectorAll('.menu-items a');
    
    menuLinks.forEach(link => {
        link.addEventListener('click', function(e) {
            // Check if link has anchor (#)
            const href = this.getAttribute('href');
            if (href.startsWith('#')) {
                e.preventDefault(); // Stop default jump behavior
                const targetSection = document.querySelector(href);
                if (targetSection) {
                    targetSection.scrollIntoView({ 
                        behavior: 'smooth',
                        block: 'start'
                    });
                }
            }
        });
    });

    // ============================================
    // 2. SCROLL REVEAL ANIMATIONS
    // ============================================
    // Elements fade in when user scrolls to them
    
    const observerOptions = {
        threshold: 0.1, // Trigger when 10% of element is visible
        rootMargin: '0px 0px -100px 0px' // Start animation 100px before element enters viewport
    };

    const observer = new IntersectionObserver(function(entries) {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.classList.add('fade-in-visible');
            }
        });
    }, observerOptions);

    // Add fade-in class to elements we want to animate
    const animatedElements = document.querySelectorAll(
        '.desc-sec .items-container, .features-sec .items, .blog-item, .try-sec .texts-container'
    );
    
    animatedElements.forEach(el => {
        el.classList.add('fade-in-hidden');
        observer.observe(el);
    });

    // ============================================
    // 3. NAVBAR SCROLL EFFECT
    // ============================================
    // Change navbar background when user scrolls down
    
    const header = document.querySelector('header');
    let lastScroll = 0;

    window.addEventListener('scroll', function() {
        const currentScroll = window.pageYOffset;

        if (currentScroll > 100) {
            header.classList.add('scrolled');
        } else {
            header.classList.remove('scrolled');
        }

        // Hide navbar on scroll down, show on scroll up
        if (currentScroll > lastScroll && currentScroll > 500) {
            header.classList.add('hidden');
        } else {
            header.classList.remove('hidden');
        }

        lastScroll = currentScroll;
    });

    // ============================================
    // 4. COUNTER ANIMATION FOR FEATURES
    // ============================================
    // Animate numbers counting up (you can add data-count attribute to elements)
    
    function animateCounter(element, target, duration = 2000) {
        let current = 0;
        const increment = target / (duration / 16); // 60fps
        
        const timer = setInterval(() => {
            current += increment;
            if (current >= target) {
                element.textContent = target;
                clearInterval(timer);
            } else {
                element.textContent = Math.floor(current);
            }
        }, 16);
    }

    // ============================================
    // 5. PARALLAX EFFECT FOR HERO SECTION
    // ============================================
    // Background moves slower than scroll for depth effect
    
    const heroSection = document.querySelector('.hero-sec');
    
    window.addEventListener('scroll', function() {
        const scrolled = window.pageYOffset;
        const parallaxSpeed = 0.5;
        
        if (heroSection) {
            heroSection.style.backgroundPositionY = `${70 + (scrolled * parallaxSpeed * 0.05)}%`;
        }
    });

    // ============================================
    // 6. FEATURE CARDS STAGGER ANIMATION
    // ============================================
    // Feature cards appear one after another with delay
    
    const featureItems = document.querySelectorAll('.features-sec .items');
    
    const featureObserver = new IntersectionObserver(function(entries) {
        entries.forEach((entry, index) => {
            if (entry.isIntersecting) {
                setTimeout(() => {
                    entry.target.classList.add('slide-in-bottom');
                }, index * 150); // 150ms delay between each card
            }
        });
    }, { threshold: 0.2 });

    featureItems.forEach(item => {
        featureObserver.observe(item);
    });

    // ============================================
    // 7. BLOG ITEMS HOVER SCALE EFFECT
    // ============================================
    // Add smooth hover effect to blog items (already in CSS, but can enhance with JS)
    
    const blogItems = document.querySelectorAll('.blog-item');
    
    blogItems.forEach(item => {
        item.addEventListener('mouseenter', function() {
            // You can add additional effects here
            this.style.boxShadow = '0 15px 40px rgba(0, 0, 0, 0.3)';
        });
        
        item.addEventListener('mouseleave', function() {
            this.style.boxShadow = 'none';
        });
    });

    // ============================================
    // 9. SCROLL PROGRESS INDICATOR
    // ============================================
    // Show how much of page user has scrolled
    
    const progressBar = document.createElement('div');
    progressBar.classList.add('scroll-progress');
    document.body.appendChild(progressBar);
    
    window.addEventListener('scroll', function() {
        const windowHeight = document.documentElement.scrollHeight - document.documentElement.clientHeight;
        const scrolled = (window.pageYOffset / windowHeight) * 100;
        progressBar.style.width = scrolled + '%';
    });

    // ============================================
    // 10. TYPING EFFECT FOR HERO TITLE (OPTIONAL)
    // ============================================
    // Make the hero title appear letter by letter
    
    const heroTitle = document.querySelector('.hero-sec h2');
    if (heroTitle) {
        const originalText = heroTitle.textContent;
        heroTitle.textContent = '';
        let charIndex = 0;
        
        function typeWriter() {
            if (charIndex < originalText.length) {
                heroTitle.textContent += originalText.charAt(charIndex);
                charIndex++;
                setTimeout(typeWriter, 80); // 80ms per character
            }
        }
        
        // Start typing effect after 500ms delay
        setTimeout(typeWriter, 500);
    }

    // ============================================
    // 11. LAZY LOADING FOR IMAGES
    // ============================================
    // Load images only when they're about to enter viewport
    
    const images = document.querySelectorAll('img[data-src]');
    
    const imageObserver = new IntersectionObserver(function(entries) {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                const img = entry.target;
                img.src = img.dataset.src;
                img.removeAttribute('data-src');
                imageObserver.unobserve(img);
            }
        });
    });
    
    images.forEach(img => imageObserver.observe(img));

    // ============================================
    // 13. FORM VALIDATION (For Sign Up Button)
    // ============================================
    // You can expand this when you add a signup form
    
    const signupBtn = document.querySelector('.try-sec button');
    if (signupBtn) {
        signupBtn.addEventListener('click', function(e) {
            e.preventDefault();
            // Add your signup logic here
            console.log('Sign up clicked - redirect to signup page');
            // window.location.href = '/signup';
        });
    }

    console.log('✅ Tasktop JS loaded successfully!');
});

document.querySelectorAll('.features-sec .items-container .items').forEach((card, index) => {
    // Create glow orb element
    const orb = document.createElement('div');
    orb.classList.add('glow-orb');
    card.prepend(orb);

    // Create particle container
    const particlesContainer = document.createElement('div');
    particlesContainer.classList.add('particles-container');
    card.prepend(particlesContainer);

    // Create particles
    for (let i = 0; i < 6; i++) {
        const particle = document.createElement('div');
        particle.classList.add('feature-particle');
        particle.style.setProperty('--delay', `${i * 0.15}s`);
        particle.style.setProperty('--x', `${Math.random() * 80 + 10}%`);
        particle.style.setProperty('--y', `${Math.random() * 80 + 10}%`);
        particle.style.setProperty('--size', `${Math.random() * 4 + 2}px`);
        particlesContainer.appendChild(particle);
    }

    // 3D tilt effect on mouse move
    card.addEventListener('mousemove', (e) => {
        const rect = card.getBoundingClientRect();
        const x = e.clientX - rect.left;
        const y = e.clientY - rect.top;
        const centerX = rect.width / 2;
        const centerY = rect.height / 2;

        const rotateX = ((y - centerY) / centerY) * -12;
        const rotateY = ((x - centerX) / centerX) * 12;

        card.style.transform = `perspective(800px) rotateX(${rotateX}deg) rotateY(${rotateY}deg) scale(1.08)`;

        // Move the glow orb to follow the cursor
        const percentX = (x / rect.width) * 100;
        const percentY = (y / rect.height) * 100;
        orb.style.background = `radial-gradient(circle at ${percentX}% ${percentY}%, rgba(80,70,229,0.5) 0%, transparent 70%)`;
    });

    card.addEventListener('mouseleave', () => {
        card.style.transform = 'perspective(800px) rotateX(0deg) rotateY(0deg) scale(1)';
        orb.style.background = 'transparent';
    });
});
