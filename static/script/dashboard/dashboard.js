// ============================================
// GLOBAL STATE & INITIALIZATION
// ============================================

const state = {
    tasks: [],
    notes: [],
    currentMonth: new Date(),
    currentView: 'dashboard'
};

// Initialize on DOM load
document.addEventListener('DOMContentLoaded', function() {
    initializeApp();
});

function initializeApp() {
    loadSavedData();
    setupEventListeners();
    updateCurrentDate();
    generateCalendar();
    updateProgress();
    updateStats();
    renderTasks();
    renderNotes();
    
    // Add gradient definition to progress ring
    addProgressGradient();
}

// ============================================
// EVENT LISTENERS
// ============================================

function setupEventListeners() {
    // Navigation
    document.querySelectorAll('.nav-item').forEach(item => {
        item.addEventListener('click', handleNavigation);
    });

    // Profile dropdown
    const profileDropdown = document.getElementById('profile-dropdown');
    const profileBtn = profileDropdown.querySelector('.profile-btn');
    
    profileBtn.addEventListener('click', (e) => {
        e.stopPropagation();
        profileDropdown.classList.toggle('active');
    });

    document.addEventListener('click', () => {
        profileDropdown.classList.remove('active');
    });

    // Calendar navigation
    document.getElementById('prev-month').addEventListener('click', () => {
        state.currentMonth.setMonth(state.currentMonth.getMonth() - 1);
        generateCalendar();
    });

    document.getElementById('next-month').addEventListener('click', () => {
        state.currentMonth.setMonth(state.currentMonth.getMonth() + 1);
        generateCalendar();
    });

    // Add task button
    document.getElementById('add-task-btn').addEventListener('click', addNewTask);
    document.getElementById('add-daily-task')?.addEventListener('click', addNewTask);
    document.getElementById('add-monthly-goal')?.addEventListener('click', addNewGoal);

    // Note form
    document.getElementById('note-form').addEventListener('submit', handleNoteSubmit);

    // AI chat
    document.getElementById('ai-chat-form').addEventListener('submit', handleAIChat);

    // Progress modal
    document.getElementById('close-progress-modal').addEventListener('click', closeProgressModal);
    document.getElementById('cancel-progress').addEventListener('click', closeProgressModal);
    document.getElementById('save-progress').addEventListener('click', saveProgress);

    // Close modal on outside click
    document.getElementById('progress-modal').addEventListener('click', (e) => {
        if (e.target.id === 'progress-modal') {
            closeProgressModal();
        }
    });
}

// ============================================
// NAVIGATION
// ============================================

function handleNavigation(e) {
    e.preventDefault();
    
    // Remove active from all nav items
    document.querySelectorAll('.nav-item').forEach(item => {
        item.classList.remove('active');
    });
    
    // Add active to clicked item
    this.classList.add('active');
    
    // Get view name
    const viewName = this.getAttribute('data-view');
    
    // Update view
    switchView(viewName);
}

function switchView(viewName) {
    // Hide all views
    document.querySelectorAll('.view').forEach(view => {
        view.classList.remove('active');
    });
    
    // Show selected view
    const targetView = document.getElementById(`${viewName}-view`);
    if (targetView) {
        targetView.classList.add('active');
    }
    
    // Update page title
    const titles = {
        'dashboard': 'Dashboard',
        'daily': 'Daily Tasks',
        'monthly': 'Monthly Goals',
        'yearly': 'Yearly Plan',
        'notes': 'My Notes',
        'ai': 'AI Assistant'
    };
    
    document.getElementById('page-title').textContent = titles[viewName] || 'Dashboard';
    state.currentView = viewName;
}

// ============================================
// DATE & CALENDAR
// ============================================

function updateCurrentDate() {
    const now = new Date();
    const options = { 
        weekday: 'long', 
        year: 'numeric', 
        month: 'long', 
        day: 'numeric' 
    };
    
    const dateStr = now.toLocaleDateString('en-US', options);
    document.getElementById('current-date').textContent = dateStr;
}

function generateCalendar() {
    const year = state.currentMonth.getFullYear();
    const month = state.currentMonth.getMonth();
    
    // Update calendar title
    const monthNames = ["January", "February", "March", "April", "May", "June",
        "July", "August", "September", "October", "November", "December"];
    document.getElementById('calendar-month-text').textContent = `${monthNames[month]} ${year}`;
    
    // Get first day and number of days
    const firstDay = new Date(year, month, 1).getDay();
    const daysInMonth = new Date(year, month + 1, 0).getDate();
    const daysInPrevMonth = new Date(year, month, 0).getDate();
    
    const calendarDays = document.getElementById('calendar-days');
    calendarDays.innerHTML = '';
    
    const today = new Date();
    const isCurrentMonth = today.getMonth() === month && today.getFullYear() === year;
    
    // Previous month days
    for (let i = firstDay - 1; i >= 0; i--) {
        const day = daysInPrevMonth - i;
        const dayElement = createDayElement(day, true, false);
        calendarDays.appendChild(dayElement);
    }
    
    // Current month days
    for (let day = 1; day <= daysInMonth; day++) {
        const isToday = isCurrentMonth && day === today.getDate();
        const hasTasks = hasTasksOnDay(year, month, day);
        const dayElement = createDayElement(day, false, isToday, hasTasks);
        calendarDays.appendChild(dayElement);
    }
    
    // Next month days
    const totalCells = calendarDays.children.length;
    const remainingCells = 42 - totalCells; // 6 rows × 7 days
    
    for (let day = 1; day <= remainingCells; day++) {
        const dayElement = createDayElement(day, true, false);
        calendarDays.appendChild(dayElement);
    }
}

function createDayElement(day, isOtherMonth, isToday, hasTasks = false) {
    const dayElement = document.createElement('div');
    dayElement.className = 'calendar-day';
    dayElement.textContent = day;
    
    if (isOtherMonth) {
        dayElement.classList.add('other-month');
    }
    
    if (isToday) {
        dayElement.classList.add('today');
    }
    
    if (hasTasks) {
        dayElement.classList.add('has-tasks');
    }
    
    return dayElement;
}

function hasTasksOnDay(year, month, day) {
    // Check if any tasks exist for this day
    const dateStr = `${year}-${String(month + 1).padStart(2, '0')}-${String(day).padStart(2, '0')}`;
    return state.tasks.some(task => task.date === dateStr);
}

// ============================================
// TASKS MANAGEMENT
// ============================================

function addNewTask() {
    const title = prompt('Enter task title:');
    if (!title) return;
    
    const priority = prompt('Enter priority (high/medium/low):', 'medium');
    const date = new Date().toISOString().split('T')[0];
    
    const task = {
        id: Date.now(),
        title: escapeHtml(title),
        priority: priority || 'medium',
        date: date,
        completed: false,
        type: 'daily'
    };
    
    state.tasks.push(task);
    saveData();
    renderTasks();
    updateProgress();
    updateStats();
}

function renderTasks() {
    const tasksList = document.getElementById('tasks-list');
    const dailyTasksList = document.getElementById('daily-tasks-list');
    
    if (!tasksList) return;
    
    const today = new Date().toISOString().split('T')[0];
    const todayTasks = state.tasks.filter(task => task.date === today);
    
    // Render dashboard tasks
    tasksList.innerHTML = todayTasks.length === 0 
        ? '<p style="color: var(--text-secondary); text-align: center; padding: 2rem;">No tasks for today. Add one to get started!</p>'
        : '';
    
    todayTasks.forEach(task => {
        tasksList.appendChild(createTaskElement(task));
    });
    
    // Render daily view tasks
    if (dailyTasksList) {
        dailyTasksList.innerHTML = state.tasks.length === 0
            ? '<p style="color: var(--text-secondary); text-align: center; padding: 2rem;">No tasks yet.</p>'
            : '';
        
        state.tasks.forEach(task => {
            dailyTasksList.appendChild(createTaskElement(task));
        });
    }
}

function createTaskElement(task) {
    const taskItem = document.createElement('div');
    taskItem.className = 'task-item';
    taskItem.innerHTML = `
        <input type="checkbox" class="task-checkbox" ${task.completed ? 'checked' : ''} data-task-id="${task.id}">
        <div class="task-content">
            <div class="task-title">${task.title}</div>
            <div class="task-meta">
                <span class="task-priority priority-${task.priority}">${task.priority.toUpperCase()}</span>
                <span>📅 ${formatDate(task.date)}</span>
            </div>
        </div>
    `;
    
    const checkbox = taskItem.querySelector('.task-checkbox');
    checkbox.addEventListener('change', function() {
        toggleTaskComplete(task.id);
    });
    
    return taskItem;
}

function toggleTaskComplete(taskId) {
    const task = state.tasks.find(t => t.id === taskId);
    if (task) {
        task.completed = !task.completed;
        saveData();
        renderTasks();
        updateProgress();
        updateStats();
    }
}

function formatDate(dateStr) {
    const date = new Date(dateStr);
    return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' });
}

// ============================================
// PROGRESS TRACKING
// ============================================

function updateProgress() {
    const today = new Date().toISOString().split('T')[0];
    const todayTasks = state.tasks.filter(task => task.date === today);
    const completedTasks = todayTasks.filter(task => task.completed);
    
    const total = todayTasks.length;
    const completed = completedTasks.length;
    const percentage = total > 0 ? Math.round((completed / total) * 100) : 0;
    
    // Update progress ring
    const circle = document.getElementById('progress-circle');
    const radius = 60;
    const circumference = 2 * Math.PI * radius;
    const offset = circumference - (percentage / 100) * circumference;
    
    circle.style.strokeDasharray = `${circumference} ${circumference}`;
    circle.style.strokeDashoffset = offset;
    
    // Update text
    document.getElementById('progress-percentage').textContent = `${percentage}%`;
    document.getElementById('completed-tasks').textContent = completed;
    document.getElementById('total-tasks').textContent = total;
}

function addProgressGradient() {
    const svg = document.querySelector('.progress-ring');
    const defs = document.createElementNS('http://www.w3.org/2000/svg', 'defs');
    const gradient = document.createElementNS('http://www.w3.org/2000/svg', 'linearGradient');
    
    gradient.setAttribute('id', 'progressGradient');
    gradient.setAttribute('x1', '0%');
    gradient.setAttribute('y1', '0%');
    gradient.setAttribute('x2', '100%');
    gradient.setAttribute('y2', '100%');
    
    const stop1 = document.createElementNS('http://www.w3.org/2000/svg', 'stop');
    stop1.setAttribute('offset', '0%');
    stop1.setAttribute('style', 'stop-color:#5046e5;stop-opacity:1');
    
    const stop2 = document.createElementNS('http://www.w3.org/2000/svg', 'stop');
    stop2.setAttribute('offset', '100%');
    stop2.setAttribute('style', 'stop-color:#7189FF;stop-opacity:1');
    
    gradient.appendChild(stop1);
    gradient.appendChild(stop2);
    defs.appendChild(gradient);
    svg.insertBefore(defs, svg.firstChild);
}

function updateStats() {
    const total = state.tasks.length;
    const completed = state.tasks.filter(task => task.completed).length;
    const pending = total - completed;
    
    document.getElementById('total-tasks-count').textContent = total;
    document.getElementById('completed-tasks-count').textContent = completed;
    document.getElementById('pending-tasks-count').textContent = pending;
}

// ============================================
// MONTHLY GOALS
// ============================================

function addNewGoal() {
    const title = prompt('Enter goal title:');
    if (!title) return;
    
    const goalsList = document.getElementById('monthly-goals-list');
    
    const goalItem = document.createElement('div');
    goalItem.className = 'goal-item';
    goalItem.innerHTML = `
        <div class="goal-header">
            <h3 class="goal-title">${escapeHtml(title)}</h3>
            <div>
                <button class="btn-secondary btn-sm decrease-progress">−</button>
                <button class="btn-primary btn-sm increase-progress">+</button>
            </div>
        </div>
        <div class="goal-progress">
            <div class="progress-bar-container">
                <div class="progress-bar" style="width: 0%"></div>
            </div>
            <div class="progress-label">Progress: <span class="progress-value">0%</span></div>
        </div>
    `;
    
    goalItem.setAttribute('data-progress', '0');
    
    // Add event listeners for progress buttons
    goalItem.querySelector('.increase-progress').addEventListener('click', () => {
        updateGoalProgress(goalItem, 10);
    });
    
    goalItem.querySelector('.decrease-progress').addEventListener('click', () => {
        updateGoalProgress(goalItem, -10);
    });
    
    goalsList.appendChild(goalItem);
}

function updateGoalProgress(goalItem, change) {
    let current = parseInt(goalItem.getAttribute('data-progress')) || 0;
    current = Math.max(0, Math.min(100, current + change));
    
    goalItem.setAttribute('data-progress', current);
    goalItem.querySelector('.progress-bar').style.width = `${current}%`;
    goalItem.querySelector('.progress-value').textContent = `${current}%`;
}

// ============================================
// NOTES MANAGEMENT
// ============================================

function handleNoteSubmit(e) {
    e.preventDefault();
    
    const title = document.getElementById('note-title').value;
    const content = document.getElementById('note-content').value;
    const category = document.getElementById('note-category').value;
    
    const note = {
        id: Date.now(),
        title: escapeHtml(title),
        content: escapeHtml(content),
        category: category,
        date: new Date().toISOString()
    };
    
    state.notes.push(note);
    saveData();
    renderNotes();
    
    // Reset form
    e.target.reset();
}

function renderNotes() {
    const notesContainer = document.getElementById('notes-container');
    if (!notesContainer) return;
    
    notesContainer.innerHTML = state.notes.length === 0
        ? '<p style="color: var(--text-secondary); text-align: center; padding: 2rem; grid-column: 1/-1;">No notes yet. Create your first note!</p>'
        : '';
    
    state.notes.forEach(note => {
        notesContainer.appendChild(createNoteElement(note));
    });
}

function createNoteElement(note) {
    const noteDiv = document.createElement('div');
    noteDiv.className = `note ${note.category}`;
    
    const timeAgo = getTimeAgo(new Date(note.date));
    
    noteDiv.innerHTML = `
        <div class="note-header">
            <h4 class="note-title">${note.title}</h4>
            <button class="note-delete" data-note-id="${note.id}">✕</button>
        </div>
        <p class="note-content">${note.content}</p>
        <div class="note-footer">
            <span class="note-category ${note.category}">${note.category}</span>
            <span>${timeAgo}</span>
        </div>
    `;
    
    noteDiv.querySelector('.note-delete').addEventListener('click', function() {
        deleteNote(note.id);
    });
    
    return noteDiv;
}

function deleteNote(noteId) {
    if (confirm('Are you sure you want to delete this note?')) {
        state.notes = state.notes.filter(note => note.id !== noteId);
        saveData();
        renderNotes();
    }
}

function getTimeAgo(date) {
    const seconds = Math.floor((new Date() - date) / 1000);
    
    const intervals = {
        year: 31536000,
        month: 2592000,
        week: 604800,
        day: 86400,
        hour: 3600,
        minute: 60
    };
    
    for (const [unit, secondsInUnit] of Object.entries(intervals)) {
        const interval = Math.floor(seconds / secondsInUnit);
        if (interval >= 1) {
            return `${interval} ${unit}${interval > 1 ? 's' : ''} ago`;
        }
    }
    
    return 'just now';
}

// ============================================
// AI ASSISTANT
// ============================================

function handleAIChat(e) {
    e.preventDefault();
    
    const input = document.getElementById('ai-input');
    const message = input.value.trim();
    
    if (!message) return;
    
    const messagesContainer = document.getElementById('ai-messages');
    
    // Add user message
    addChatMessage(messagesContainer, message, true);
    
    // Clear input
    input.value = '';
    
    // Simulate AI response
    setTimeout(() => {
        const response = generateAIResponse(message);
        addChatMessage(messagesContainer, response, false);
    }, 1000);
}

function addChatMessage(container, text, isUser) {
    const messageDiv = document.createElement('div');
    messageDiv.className = `chat-message ${isUser ? 'user' : 'ai'}`;
    
    const avatar = isUser 
        ? 'https://ui-avatars.com/api/?name=John+Doe&background=5046e5&color=fff'
        : 'https://ui-avatars.com/api/?name=AI&background=7189FF&color=fff';
    
    const now = new Date().toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit' });
    
    messageDiv.innerHTML = `
        <img src="${avatar}" alt="Avatar" class="message-avatar">
        <div class="message-content">
            <p class="message-text">${escapeHtml(text)}</p>
            <span class="message-time">${now}</span>
        </div>
    `;
    
    container.appendChild(messageDiv);
    container.scrollTop = container.scrollHeight;
}

function generateAIResponse(message) {
    const responses = {
        'hello': 'Hello! How can I help you with your tasks today?',
        'help': 'I can help you manage tasks, set priorities, and plan your schedule. What would you like to do?',
        'task': `You have ${state.tasks.length} tasks in total. ${state.tasks.filter(t => t.completed).length} are completed.`,
        'priority': 'I recommend focusing on high-priority tasks first. Would you like me to list them?',
        'default': 'I understand. Is there anything specific you need help with regarding your tasks or planning?'
    };
    
    const lowerMessage = message.toLowerCase();
    
    for (const [key, response] of Object.entries(responses)) {
        if (lowerMessage.includes(key)) {
            return response;
        }
    }
    
    return responses.default;
}

// ============================================
// PROGRESS MODAL
// ============================================

let currentTargetTask = null;

function openProgressModal(taskElement) {
    currentTargetTask = taskElement;
    document.getElementById('progress-modal').classList.add('active');
    document.getElementById('progress-input').focus();
}

function closeProgressModal() {
    document.getElementById('progress-modal').classList.remove('active');
    document.getElementById('progress-input').value = '';
    currentTargetTask = null;
}

function saveProgress() {
    const value = parseInt(document.getElementById('progress-input').value);
    
    if (currentTargetTask && value >= 0 && value <= 100) {
        currentTargetTask.setAttribute('data-progress', value);
        const progressBar = currentTargetTask.querySelector('.progress-bar');
        const progressLabel = currentTargetTask.querySelector('.progress-value');
        
        if (progressBar) progressBar.style.width = `${value}%`;
        if (progressLabel) progressLabel.textContent = `${value}%`;
    }
    
    closeProgressModal();
}

// ============================================
// DATA PERSISTENCE
// ============================================

function saveData() {
    localStorage.setItem('taskflow_tasks', JSON.stringify(state.tasks));
    localStorage.setItem('taskflow_notes', JSON.stringify(state.notes));
}

function loadSavedData() {
    const savedTasks = localStorage.getItem('taskflow_tasks');
    const savedNotes = localStorage.getItem('taskflow_notes');
    
    if (savedTasks) {
        state.tasks = JSON.parse(savedTasks);
    }
    
    if (savedNotes) {
        state.notes = JSON.parse(savedNotes);
    }
}

// ============================================
// UTILITY FUNCTIONS
// ============================================

function escapeHtml(text) {
    const map = {
        '&': '&amp;',
        '<': '&lt;',
        '>': '&gt;',
        '"': '&quot;',
        "'": '&#039;'
    };
    return text.replace(/[&<>"']/g, m => map[m]);
}


