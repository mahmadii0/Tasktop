// ============================================
// GLOBAL STATE & INITIALIZATION
// ============================================

const state = {
    tasks: [],
    notes: [],
    currentMonth: new Date(),
    currentView: 'dashboard',
    selectedDate: new Date()
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
    const selectedView = document.getElementById(`${viewName}-view`);
    if (selectedView) {
        selectedView.classList.add('active');
    }

    state.currentView = viewName;
}

// ============================================
// DATE & TIME
// ============================================

function updateCurrentDate() {
    const dateElement = document.getElementById('current-date');
    if (!dateElement) return;

    const options = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' };
    const currentDate = new Date().toLocaleDateString('en-US', options);
    dateElement.textContent = currentDate;
}


// ============================================
// CALENDAR WIDGET
// ============================================

function generateCalendar() {
    const calendarGrid = document.getElementById('calendar-grid');
    const currentMonthElement = document.getElementById('current-month');

    if (!calendarGrid) return;

    const year = state.currentMonth.getFullYear();
    const month = state.currentMonth.getMonth();

    // Update month display
    const monthNames = ['January', 'February', 'March', 'April', 'May', 'June',
        'July', 'August', 'September', 'October', 'November', 'December'];
    currentMonthElement.textContent = `${monthNames[month]} ${year}`;

    // Clear calendar
    calendarGrid.innerHTML = '';

    // Get first day of month and number of days
    const firstDay = new Date(year, month, 1).getDay();
    const daysInMonth = new Date(year, month + 1, 0).getDate();

    // Add empty cells for days before month starts
    for (let i = 0; i < firstDay; i++) {
        const emptyDay = document.createElement('div');
        emptyDay.className = 'calendar-day empty';
        calendarGrid.appendChild(emptyDay);
    }

    // Add days of month
    const today = new Date();
    for (let day = 1; day <= daysInMonth; day++) {
        const dayElement = document.createElement('div');
        dayElement.className = 'calendar-day';
        dayElement.textContent = day;

        // Highlight today
        if (year === today.getFullYear() &&
            month === today.getMonth() &&
            day === today.getDate()) {
            dayElement.classList.add('today');
        }

        // Check if there are tasks on this day
        const dateStr = `${year}-${String(month + 1).padStart(2, '0')}-${String(day).padStart(2, '0')}`;
        const hasTasks = state.tasks.some(task => task.date === dateStr);

        if (hasTasks) {
            dayElement.classList.add('has-tasks');
        }

        calendarGrid.appendChild(dayElement);
    }
}

// ============================================
// TASKS
// ============================================

function renderTasks() {
    const tasksList = document.getElementById('tasks-list');
    if (!tasksList) return;

    // Try to fetch tasks from server endpoint, fall back to local state on error
    fetch('/dashboard/daily-goals')
        .then(res => {
            if (!res.ok) throw new Error('Network response was not ok');
            return res.json();
        })
        .then(data => {
            // Accept either an array or an object with a `tasks` property
            const tasksFromServer = Array.isArray(data) ? data : (data.tasks || []);
            if (tasksFromServer && tasksFromServer.length) {
                state.tasks = tasksFromServer;
                // Persist fetched tasks locally so UI remains consistent
                saveData();
            }
        })
        .catch(err => {
            console.error('Failed to fetch tasks from /dashboard/daily-goals, using local data:', err);
        })
        .finally(() => {
            renderTasksFromState();
        });
}

function renderTasksFromState() {
    const tasksList = document.getElementById('tasks-list');
    const dailyTasksList = document.getElementById('daily-tasks-list');
    if (!tasksList) return;

    const today = new Date().toISOString().split('T')[0];
    const todayTasks = state.tasks.filter(task => task.date === today);

    // Render dashboard tasks (always show today)
    tasksList.innerHTML = todayTasks.length === 0
        ? '<p style="color: var(--text-secondary); text-align: center; padding: 2rem;">No tasks for today. Add one to get started!</p>'
        : '';

    todayTasks.forEach(task => {
        tasksList.appendChild(createTaskElement(task));
    });

    // Render daily view tasks based on selected date
    if (dailyTasksList) {
        if (typeof renderTasksForDate === 'function') {
            renderTasksForDate(state.selectedDate);
        }
    }
}

function createTaskElement(task) {
    const taskElement = document.createElement('div');
    taskElement.className = `task-item ${task.completed ? 'completed' : ''}`;
    taskElement.setAttribute('data-task-id', task.id);

    const priorityClass = `priority-${task.priority}`;
    const priorityText = task.priority.charAt(0).toUpperCase() + task.priority.slice(1);

    taskElement.innerHTML = `
        <div class="task-checkbox">
            <input type="checkbox" ${task.completed ? 'checked' : ''} 
                   onchange="toggleTask('${task.id}')">
        </div>
        <div class="task-details">
            <div class="task-content">
                <h4 class="task-title">${escapeHtml(task.title)}</h4>
                ${task.description ? `<p class="task-description">${escapeHtml(task.description)}</p>` : ''}
            </div>
            <div class="task-meta">
                <span class="task-priority ${priorityClass}">${priorityText}</span>
                <span class="task-date">
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect>
                        <line x1="16" y1="2" x2="16" y2="6"></line>
                        <line x1="8" y1="2" x2="8" y2="6"></line>
                        <line x1="3" y1="10" x2="21" y2="10"></line>
                    </svg>
                    ${formatTaskDate(task.date)}
                </span>
            </div>
        </div>
        <div class="task-actions">
            <button class="task-action-btn" onclick="deleteTask('${task.id}')" title="Delete">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <polyline points="3 6 5 6 21 6"></polyline>
                    <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path>
                </svg>
            </button>
        </div>
    `;

    return taskElement;
}

function formatTaskDate(dateStr) {
    const date = new Date(dateStr);
    const today = new Date();
    const tomorrow = new Date(today);
    tomorrow.setDate(tomorrow.getDate() + 1);

    if (dateStr === today.toISOString().split('T')[0]) {
        return 'Today';
    } else if (dateStr === tomorrow.toISOString().split('T')[0]) {
        return 'Tomorrow';
    } else {
        const options = { month: 'short', day: 'numeric' };
        return date.toLocaleDateString('en-US', options);
    }
}

function addNewTask() {
    const title = prompt('Enter task title:');
    if (!title) return;

    const description = prompt('Enter task description (optional):') || '';
    const priority = prompt('Enter priority (low/medium/high):', 'medium') || 'medium';
    const dateInput = prompt('Enter date (YYYY-MM-DD):', new Date().toISOString().split('T')[0]);

    const task = {
        id: Date.now().toString(),
        title,
        description,
        priority: ['low', 'medium', 'high'].includes(priority.toLowerCase()) ? priority.toLowerCase() : 'medium',
        date: dateInput || new Date().toISOString().split('T')[0],
        completed: false
    };

    state.tasks.push(task);
    saveData();
    renderTasks();
    updateStats();
    updateProgress();
}

function toggleTask(taskId) {
    const task = state.tasks.find(t => t.id === taskId);
    if (task) {
        task.completed = !task.completed;
        saveData();
        renderTasks();
        updateStats();
        updateProgress();
    }
}

function deleteTask(taskId) {
    if (confirm('Are you sure you want to delete this task?')) {
        state.tasks = state.tasks.filter(t => t.id !== taskId);
        saveData();
        renderTasks();
        updateStats();
        updateProgress();
    }
}

// ============================================
// STATS
// ============================================

function updateStats() {
    const totalTasks = state.tasks.length;
    const completedTasks = state.tasks.filter(t => t.completed).length;
    const pendingTasks = totalTasks - completedTasks;

    document.getElementById('total-tasks').textContent = totalTasks;
    document.getElementById('completed-tasks').textContent = completedTasks;
    document.getElementById('pending-tasks').textContent = pendingTasks;
}

// ============================================
// PROGRESS
// ============================================

function updateProgress() {
    const totalTasks = state.tasks.length;
    const completedTasks = state.tasks.filter(t => t.completed).length;
    const percentage = totalTasks > 0 ? Math.round((completedTasks / totalTasks) * 100) : 0;

    document.getElementById('progress-percentage').textContent = `${percentage}%`;
    document.getElementById('completed-count').textContent = completedTasks;
    document.getElementById('total-count').textContent = totalTasks;

    // Update progress ring
    const circle = document.querySelector('.progress-ring-fill');
    if (circle) {
        const radius = circle.r.baseVal.value;
        const circumference = 2 * Math.PI * radius;
        const offset = circumference - (percentage / 100) * circumference;
        circle.style.strokeDasharray = `${circumference} ${circumference}`;
        circle.style.strokeDashoffset = offset;
    }
}

function addProgressGradient() {
    const svg = document.querySelector('.progress-ring');
    if (!svg || document.getElementById('progressGradient')) return;

    const defs = document.createElementNS('http://www.w3.org/2000/svg', 'defs');
    const gradient = document.createElementNS('http://www.w3.org/2000/svg', 'linearGradient');
    gradient.setAttribute('id', 'progressGradient');
    gradient.innerHTML = `
        <stop offset="0%" stop-color="#5046e5" />
        <stop offset="100%" stop-color="#7189FF" />
    `;
    defs.appendChild(gradient);
    svg.insertBefore(defs, svg.firstChild);
}

// ============================================
// NOTES
// ============================================

function renderNotes() {
    const notesList = document.getElementById('notes-list');
    if (!notesList) return;

    notesList.innerHTML = state.notes.length === 0
        ? '<p style="color: var(--text-secondary); text-align: center; padding: 2rem;">No notes yet. Create your first note!</p>'
        : '';

    state.notes.forEach(note => {
        const noteElement = document.createElement('div');
        noteElement.className = 'note-item';
        noteElement.innerHTML = `
            <h4 class="note-title">${escapeHtml(note.title)}</h4>
            <p class="note-content">${escapeHtml(note.content)}</p>
            <div class="note-meta">
                <span class="note-date">${new Date(note.date).toLocaleDateString()}</span>
                <button class="note-delete" onclick="deleteNote('${note.id}')">
                    <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polyline points="3 6 5 6 21 6"></polyline>
                        <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path>
                    </svg>
                </button>
            </div>
        `;
        notesList.appendChild(noteElement);
    });
}

function handleNoteSubmit(e) {
    e.preventDefault();

    const title = document.getElementById('note-title').value;
    const content = document.getElementById('note-content').value;

    if (!title || !content) return;

    const note = {
        id: Date.now().toString(),
        title,
        content,
        date: new Date().toISOString()
    };

    state.notes.unshift(note);
    saveData();
    renderNotes();

    // Clear form
    document.getElementById('note-title').value = '';
    document.getElementById('note-content').value = '';
}

function deleteNote(noteId) {
    if (confirm('Are you sure you want to delete this note?')) {
        state.notes = state.notes.filter(n => n.id !== noteId);
        saveData();
        renderNotes();
    }
}

// ============================================
// AI ASSISTANT
// ============================================

function handleAIChat(e) {
    e.preventDefault();

    const input = document.getElementById('ai-input');
    const message = input.value.trim();

    if (!message) return;

    // Add user message
    addChatMessage(message, 'user');

    // Clear input
    input.value = '';

    // Simulate AI response
    setTimeout(() => {
        const response = generateAIResponse(message);
        addChatMessage(response, 'ai');
    }, 1000);
}

function addChatMessage(text, sender) {
    const container = document.getElementById('chat-messages');
    const messageDiv = document.createElement('div');
    messageDiv.className = `chat-message ${sender}`;

    const avatar = sender === 'user'
        ? 'https://api.dicebear.com/7.x/avataaars/svg?seed=User'
        : 'https://api.dicebear.com/7.x/bottts/svg?seed=AI';

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
// MONTHLY GOALS (Placeholder)
// ============================================

function addNewGoal() {
    alert('Monthly goals feature coming soon!');
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
