 // Profile dropdown toggle
        document.getElementById('profile-dropdown').addEventListener('click', function() {
            const dropdown = this.querySelector('.dropdown');
            dropdown.classList.toggle('show');
        });

        // Close dropdown when clicking outside
        window.addEventListener('click', function(event) {
            if (!event.target.closest('#profile-dropdown')) {
                const dropdown = document.querySelector('#profile-dropdown .dropdown');
                if (dropdown.classList.contains('show')) {
                    dropdown.classList.remove('show');
                }
            }
        });
        document.addEventListener('DOMContentLoaded', function() {
            // Set current date
            const now = new Date();
            const options = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' };
            document.getElementById('current-date').textContent = now.toLocaleDateString('en-US', options);

            
            // View tabs functionality
            const viewTabs = document.querySelectorAll('.view-tab');
            const viewContents = document.querySelectorAll('.view-content');

            viewTabs.forEach(tab => {
                tab.addEventListener('click', () => {
                    // Remove active class from all tabs
                    viewTabs.forEach(t => t.classList.remove('tab-active'));
                    // Add active class to clicked tab
                    tab.classList.add('tab-active');
                    
                    // Hide all content
                    viewContents.forEach(content => content.classList.add('hidden'));
                    // Show content corresponding to clicked tab
                    const view = tab.getAttribute('data-view');
                    document.getElementById(`${view}-view`).classList.remove('hidden');
                });
            })});

            // Task checkbox functionality
           document.addEventListener("DOMContentLoaded", function () {
             const tabButtons = document.querySelectorAll('.tab-button');
             const views = document.querySelectorAll('.view-content');

           tabButtons.forEach(button => {
             button.addEventListener('click', () => {
               const targetView = button.getAttribute('data-view');
 
            views.forEach(view => view.classList.add('hidden'));

             const activeView = document.getElementById(targetView);
               if (activeView) {
                 activeView.classList.remove('hidden');
         }
      });
    });
  });

            // Task checkbox functionality from header
            document.querySelectorAll('a[data-view]').forEach(button => {
                    button.addEventListener('click', (e) => {
                        e.preventDefault();

                        const targetView = button.getAttribute('data-view');
                        const views = ['daily', 'monthly', 'yearly'];

                        views.forEach(view => {
                            const el = document.getElementById(`${view}-view`);
                            if (el) {
                                if (view === targetView) {
                                    el.classList.remove('hidden');
                                } else {
                                    el.classList.add('hidden');
                                }
                            }
                        });
                    });
                });

            // Task checkbox functionality from dashboard
            let currentMainView = 'daily-view';
            document.querySelectorAll('.sidebar-item').forEach(link => {
                link.addEventListener('click', function (e) {
                    e.preventDefault();
                    const targetId = this.getAttribute('href').substring(1);
                    if (targetId === 'Note') {
                        document.querySelectorAll('.view-content').forEach(view => {
                            if (view.id !== 'Note' && view.id !== currentMainView) {
                                view.classList.add('hidden');
                            }
                        });
                        const noteEl = document.getElementById('Note');
                        noteEl?.classList.remove('hidden');
                        setTimeout(() => {
                            noteEl.scrollIntoView({ behavior: 'smooth', block: 'start' });
                        }, 100);
                    } else {
                        currentMainView = targetId;
                        document.querySelectorAll('.view-content').forEach(view => {
                            if (view.id !== 'Note') {
                                view.classList.add('hidden');
                            }
                        });
                        document.getElementById(targetId)?.classList.remove('hidden');
                    }

                    document.querySelectorAll('.sidebar-item').forEach(item => {
                        item.classList.remove('active', 'bg-gray-200', 'text-indigo-700');
                    });
                    this.classList.add('active', 'bg-gray-200', 'text-indigo-700');
                });
            });


   // Calendar functionality
   function generateCalendar() {
                const calendarDays = document.getElementById('calendar-days');
                calendarDays.innerHTML = '';
                
                const currentDate = new Date();
                const month = currentDate.getMonth();
                const year = currentDate.getFullYear();
                
                document.getElementById('calendar-month').textContent = new Date(year, month).toLocaleDateString('en-US', { month: 'long', year: 'numeric' });
                
                const firstDay = new Date(year, month, 1).getDay();
                const daysInMonth = new Date(year, month + 1, 0).getDate();
                
                // Add empty cells for days before the first day of the month
                for (let i = 0; i < firstDay; i++) {
                    const emptyDay = document.createElement('div');
                    emptyDay.className = 'calendar-day-disabled h-8 flex items-center justify-center text-gray-300';
                    calendarDays.appendChild(emptyDay);
                }
                
                // Add cells for each day of the month
                for (let i = 1; i <= daysInMonth; i++) {
                    const day = document.createElement('div');
                    day.className = 'calendar-day h-8 flex items-center justify-center text-sm rounded-full cursor-pointer';
                    
                    // Highlight current day
                    if (i === currentDate.getDate()) {
                        day.classList.add('calendar-day-selected');
                    }
                    
                    // Add task indicator for some random days
                    if ([3, 7, 12, 18, 25].includes(i)) {
                        day.classList.add('calendar-day-has-task');
                    }
                    
                    day.textContent = i;
                    day.addEventListener('click', function() {
                        document.querySelectorAll('.calendar-day').forEach(d => d.classList.remove('calendar-day-selected'));
                        this.classList.add('calendar-day-selected');
                    });
                    
                    calendarDays.appendChild(day);
                }
            }
            
            generateCalendar();
            
            document.getElementById('prev-month').addEventListener('click', function() {
                // In a real app, this would show the previous month
                alert('Navigate to previous month');
            });
            
            document.getElementById('next-month').addEventListener('click', function() {
                // In a real app, this would show the next month
                alert('Navigate to next month');
            });

            // Update progress indicators
            function updateProgress() {
                const totalTasks = document.querySelectorAll('.task-checkbox').length;
                const completedTasks = document.querySelectorAll('.task-checkbox:checked').length;
                const percentage = totalTasks > 0 ? Math.round((completedTasks / totalTasks) * 100) : 0;
                
                document.getElementById('progress-percentage').textContent = `${percentage}%`;
                
                // Update the progress ring
                const circle = document.getElementById('progress-circle');
                const radius = 16;
                const circumference = 2 * Math.PI * radius;
                const offset = circumference - (percentage / 100) * circumference;
                circle.style.strokeDasharray = `${circumference} ${circumference}`;
                circle.style.strokeDashoffset = offset;
            }
            
            updateProgress();

        // Tab switching
        document.querySelectorAll('.tab-button').forEach(tab => {
            tab.addEventListener('click', function() {
                document.querySelectorAll('.tab-button').forEach(t => {
                    t.classList.remove('active');
                    t.classList.add('text-gray-600');
                });
                this.classList.add('active');
                this.classList.remove('text-gray-600');
                updateViewForTab(this.textContent.trim());
            });
        });

        // Sidebar navigation
        document.querySelectorAll('.sidebar-item').forEach(item => {
            item.addEventListener('click', function(e) {
                e.preventDefault();
                document.querySelectorAll('.sidebar-item').forEach(i => {
                    i.classList.remove('active');
                });
                this.classList.add('active');
                updateContentForNavItem(this.textContent.trim());
            });
        });

        // Header navigation
        document.querySelectorAll('header nav a').forEach(link => {
            link.addEventListener('click', function(e) {
                e.preventDefault();
                document.querySelectorAll('header nav a').forEach(l => {
                    l.classList.remove('active-nav');
                });
                this.classList.add('active-nav');
                updateContentForNavItem(this.textContent.trim());
            });
        });

        // Task checkbox functionality
        document.querySelectorAll('.task-item input[type="checkbox"]').forEach(checkbox => {
            checkbox.addEventListener('change', function() {
                const taskText = this.nextElementSibling;
                if (this.checked) {
                    taskText.style.textDecoration = 'line-through';
                    taskText.style.color = '#9ca3af';
                    updateProgress();
                } else {
                    taskText.style.textDecoration = 'none';
                    taskText.style.color = '#1f2937';
                    updateProgress();
                }
            });
        });

        // AI Assistant functionality
        const aiInput = document.querySelector('input[placeholder="Ask your AI assistant..."]');
        const aiSendButton = aiInput.nextElementSibling;
        const aiChatContainer = document.querySelector('.ai-message').parentElement;

        aiSendButton.addEventListener('click', function() {
            sendMessageToAI();
        });

        aiInput.addEventListener('keypress', function(e) {
            if (e.key === 'Enter') {
                sendMessageToAI();
            }
        });

        function sendMessageToAI() {
            const message = aiInput.value.trim();
            if (message) {
                // Add user message
                const userMessageDiv = document.createElement('div');
                userMessageDiv.className = 'user-message p-3 rounded-lg';
                userMessageDiv.innerHTML = `<p class="text-sm text-gray-700">${message}</p>`;
                aiChatContainer.appendChild(userMessageDiv);
                
                // Clear input
                aiInput.value = '';
                
                // Scroll to bottom
                aiChatContainer.scrollTop = aiChatContainer.scrollHeight;
                
                // Simulate AI response after a short delay
                setTimeout(() => {
                    const aiResponse = getAIResponse(message);
                    const aiMessageDiv = document.createElement('div');
                    aiMessageDiv.className = 'ai-message p-3 rounded-lg';
                    aiMessageDiv.innerHTML = `<p class="text-sm text-gray-700">${aiResponse}</p>`;
                    aiChatContainer.appendChild(aiMessageDiv);
                    
                    // Scroll to bottom again
                    aiChatContainer.scrollTop = aiChatContainer.scrollHeight;
                }, 1000);
            }
        }

        function getAIResponse(message) {
            // Simple AI response logic - in a real app this would connect to an AI service
            const responses = [
                "I've analyzed your schedule and can help you optimize your day for maximum productivity.",
                "Based on your current tasks, I recommend focusing on completing the high-priority items first.",
                "I've noticed a pattern in your task completion. Would you like me to suggest a more efficient workflow?",
                "Your progress this week is impressive! You're 15% ahead of your usual pace.",
                "I can help you break down this large task into smaller, more manageable steps if you'd like."
            ];
            return responses[Math.floor(Math.random() * responses.length)];
        }

        function updateProgress() {
            const totalTasks = document.querySelectorAll('.task-item').length;
            const completedTasks = document.querySelectorAll('.task-item input[type="checkbox"]:checked').length;
            const progressPercent = totalTasks > 0 ? Math.round((completedTasks / totalTasks) * 100) : 0;
            
            // Update progress ring
            const circle = document.querySelector('.progress-ring-circle');
            const radius = circle.r.baseVal.value;
            const circumference = radius * 2 * Math.PI;
            circle.style.strokeDasharray = `${circumference} ${circumference}`;
            const offset = circumference - (progressPercent / 100) * circumference;
            circle.style.strokeDashoffset = offset;
            
            // Update percentage text
            document.querySelector('.progress-ring + div span').textContent = `${progressPercent}%`;
        }

 // Modal functionality
 const addTaskModal = document.getElementById('add-task-modal');
            const addNoteModal = document.getElementById('add-note-modal');
            
            document.getElementById('add-daily-task').addEventListener('click', function() {
                addTaskModal.classList.remove('hidden');
            });
            
            document.getElementById('add-monthly-goal').addEventListener('click', function() {
                addTaskModal.classList.remove('hidden');
            });
            
            document.getElementById('add-yearly-objective').addEventListener('click', function() {
                addTaskModal.classList.remove('hidden');
            });
            
            document.getElementById('close-task-modal').addEventListener('click', function() {
                addTaskModal.classList.add('hidden');
            });
            
            document.getElementById('cancel-task').addEventListener('click', function() {
                addTaskModal.classList.add('hidden');
            });
            
            document.getElementById('add-note').addEventListener('click', function() {
                addNoteModal.classList.remove('hidden');
            });
            
            document.getElementById('close-note-modal').addEventListener('click', function() {
                addNoteModal.classList.add('hidden');
            });
            
            document.getElementById('cancel-note').addEventListener('click', function() {
                addNoteModal.classList.add('hidden');
            });
            
            // Form submissions
            document.getElementById('task-form').addEventListener('submit', function(e) {
                e.preventDefault();
                
                const title = document.getElementById('task-title').value;
                const priority = document.getElementById('task-priority').value;
                
                //Task type
                const taskType = document.getElementById('task-type').value;
                let taskTimeLabel = '';

                if (taskType === 'daily') {
                  taskTimeLabel = document.getElementById('task-due-date').value;
                } else if (taskType === 'monthly') {
                  const monthVal = document.getElementById('task-due-month').value;
                  const monthName = new Date(2000, monthVal, 1).toLocaleString('en-US', { month: 'long' });
                  taskTimeLabel = monthName;
                } else if (taskType === 'yearly') {
                  taskTimeLabel = document.getElementById('task-due-year').value;
                }
              
                if (title) {
                    // Add new task to daily/monthly/yearly tasks
                    let taskList;
                    if (taskType === 'daily') {
                      taskList = document.getElementById('daily-tasks');
                    } else if (taskType === 'monthly') {
                      taskList = document.getElementById('monthly-goals');
                    } else if (taskType === 'yearly') {
                      taskList = document.getElementById('yearly-objectives');
                    }
 
              //
            const priorityClass = {
                'low': 'bg-blue-100 text-blue-800',
                'medium': 'bg-yellow-100 text-yellow-800',
                'high': 'bg-green-100 text-green-800'
            };

            const priorityLabel = {
                'low': 'Low',
                'medium': 'Medium',
                'high': 'High'
            };


            const taskId = `task-${Date.now()}`;

            const newTask = document.createElement('div');
            newTask.className = 'task-item flex items-center justify-between p-3 bg-white border rounded-lg mb-2';
            newTask.setAttribute('data-id', taskId); 

            if (taskType === 'daily') {
                newTask.innerHTML = `
                  <div class="flex items-center">
                      <input type="checkbox" class="task-checkbox h-5 w-5 text-indigo-600 rounded border-gray-300 focus:ring-indigo-500">
                      <span class="ml-3 text-gray-800">${title}</span>
                  </div>
                  <div class="flex items-center space-x-2">
                      <span class="text-xs ${priorityClass[priority]} px-2 py-1 rounded">${priorityLabel[priority]}</span>
                      <span class="text-xs text-gray-500">${taskTimeLabel}</span>
                      <button class="delete-task text-gray-400 hover:text-red-500" data-id="${taskId}">
                          <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4
                                    a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                          </svg>
                      </button>
                  </div>
                `;

              } else if (taskType === 'monthly') {

                newTask.className = 'task-item monthly-task bg-white border rounded-lg p-3 mb-2';
                newTask.classList.add('monthly-task');
                newTask.setAttribute('data-progress', '0');
                newTask.innerHTML = `
                  <div class="flex flex-col space-y-1 w-full">
                    <h3 class="text-md font-semibold text-indigo-700">${title}</h3>
                    <p class="text-sm text-gray-500">${taskTimeLabel}</p>
                    <div class="flex justify-between items-center">
                      <div class="w-full mr-4">
                        <div class="flex justify-between text-xs mb-1">
                          <span>Progress</span>
                          <span class="progress-value">0%</span>
                        </div>
                        <div class="w-full bg-gray-200 rounded-full h-2">
                          <div class="progress-bar bg-indigo-500 h-2 rounded-full" style="width: 0%"></div>
                        </div>
                      </div>
                      <div class="flex flex-col space-y-1 items-center ml-2">
                        <button class="increase-progress bg-indigo-100 text-indigo-600 text-xs px-2 rounded hover:bg-indigo-200">+</button>
                        <button class="decrease-progress bg-red-100 text-red-600 text-xs px-2 rounded hover:bg-red-200">−</button>
                      </div>
                    </div>
                    <div class="flex justify-end">
                      <button class="delete-task text-gray-400 hover:text-red-500 mt-2" data-id="${taskId}">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4
                            a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                        </svg>
                      </button>
                    </div>
                  </div>
                `;

              } else if (taskType === 'yearly') {

                newTask.className = 'task-item yearly-task bg-white border rounded-lg p-3 mb-2';
                newTask.classList.add('yearly-task');
                newTask.setAttribute('data-progress', '0');
                newTask.innerHTML = `
                  <div class="flex flex-col space-y-2 w-full">
                    <div class="flex justify-between items-start">
                      <h3 class="font-semibold text-gray-800 text-lg">${title}</h3>
                      <span class="text-xs ${priorityClass[priority]} px-2 py-1 rounded">${priorityLabel[priority]}</span>
                    </div>
                    <p class="text-gray-600">${taskTimeLabel}</p>
                    <div class="flex justify-between text-sm">
                      <span>Progress</span>
                      <span class="progress-value">0%</span>
                    </div>
                    <div class="w-full bg-gray-200 rounded-full h-3">
                      <div class="progress-bar bg-indigo-500 h-3 rounded-full" style="width: 0%"></div>
                    </div>
                    <div class="flex justify-between items-center mt-2">
                      <span class="text-xs text-gray-500">Started: ${new Date().toLocaleDateString()}</span>
                      <button class="update-progress text-indigo-600 hover:text-indigo-800 text-sm font-medium">Update Progress</button>
                    </div>
                    <div class="flex justify-end">
                       <button class="delete-task text-gray-400 hover:text-red-500 mt-2" data-id="${taskId}">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4
                            a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                        </svg>
                      </button>
                    </div>
                  </div>
                `;
              }

            taskList.prepend(newTask);

            // Add checkbox functionality to new task
            const newCheckbox = newTask.querySelector('.task-checkbox');
            newCheckbox.addEventListener('change', function () {
                const taskText = this.nextElementSibling;
                if (this.checked) {
                    taskText.classList.add('line-through', 'text-gray-500');
                } else {
                   taskText.classList.remove('line-through', 'text-gray-500');
                }
                updateProgress();
            });

            // Delete Task
            document.addEventListener('click', (e) => {
  if (e.target.closest('.delete-task')) {
    const button = e.target.closest('.delete-task');
    const id = button.getAttribute('data-id');
    const taskToRemove = button.closest(`[data-id="${id}"]`);
    if (taskToRemove) {
      taskToRemove.remove();
    }
  }
});
                    
                    // Reset form and close modal
                    document.getElementById('task-title').value = '';
                    document.getElementById('task-description').value = '';
                    addTaskModal.classList.add('hidden');
                    
                    // Update progress
                    updateProgress();
                }
            });
            
              // Task type chosen
              document.getElementById('task-type').addEventListener('change', function () {
                            const type = this.value;
                            document.getElementById('daily-date-picker').classList.toggle('hidden', type !== 'daily');
                            document.getElementById('monthly-date-picker').classList.toggle('hidden', type !== 'monthly');
                            document.getElementById('yearly-date-picker').classList.toggle('hidden', type !== 'yearly');
                });

                //
              let currentTargetTask = null;

              document.addEventListener('click', function (e) {

                if (e.target.closest('.update-progress')) {
                  currentTargetTask = e.target.closest('.yearly-task');
                  document.getElementById('progress-modal').classList.remove('hidden');
                }

                if (e.target.id === 'cancel-progress') {
                  document.getElementById('progress-modal').classList.add('hidden');
                  currentTargetTask = null;
                }

                if (e.target.id === 'confirm-progress' && currentTargetTask) {
                  const newProgress = parseInt(document.getElementById('progress-input').value);
              
                  if (!isNaN(newProgress) && newProgress >= 0 && newProgress <= 100) {
                    const bar = currentTargetTask.querySelector('.progress-bar');
                    const label = currentTargetTask.querySelector('.progress-value');

                    bar.style.width = `${newProgress}%`;
                    label.textContent = `${newProgress}%`;
                    currentTargetTask.setAttribute('data-progress', newProgress);
                  }

                  document.getElementById('progress-modal').classList.add('hidden');
                  document.getElementById('progress-input').value = '';
                  currentTargetTask = null;
                }
              });

              // 
              document.addEventListener('click', function (e) {
                const parent = e.target.closest('.monthly-task');
                if (!parent) return;

                const progressBar = parent.querySelector('.progress-bar');
                const progressText = parent.querySelector('.progress-value');

                let current = parseInt(parent.getAttribute('data-progress'));

                if (e.target.closest('.increase-progress')) {
                  if (current < 100) current += 10;
                } else if (e.target.closest('.decrease-progress')) {
                  if (current > 0) current -= 10;
                } else {
                  return;
                }
              
                // 
                parent.setAttribute('data-progress', current);
                progressBar.style.width = `${current}%`;
                progressText.textContent = `${current}%`;
              });

            // note 
            document.getElementById('note-form').addEventListener('submit', function(e) {
                e.preventDefault();
                
                const title = document.getElementById('note-title').value;
                const content = document.getElementById('note-content').value;
                const category = document.getElementById('note-category').value;
                
                if (title && content) {
                    // Add new note (in a real app, this would be more sophisticated)
                    const notesContainer = document.getElementById('notes-container');
                    
                    const bgColors = {
                        'work': 'bg-yellow-50 border-yellow-200',
                        'personal': 'bg-blue-50 border-blue-200',
                        'ideas': 'bg-green-50 border-green-200',
                        'other': 'bg-purple-50 border-purple-200'
                    };
                    
                    const newNote = document.createElement('div');
                    newNote.className = `note ${bgColors[category]} p-4 rounded-lg`;
                    
                    const currentDate = new Date().toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: 'numeric' });

                    newNote.innerHTML = `
                        <div class="flex justify-between items-start mb-2">
                            <h3 class="font-medium text-gray-800">${title}</h3>
                            <div class="flex space-x-1">
                                <button class="text-gray-400 hover:text-gray-600">
                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                                    </svg>
                                </button>
                                <button class="text-gray-400 hover:text-red-500">
                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                                    </svg>
                                </button>
                            </div>
                        </div>
                        <p class="text-sm text-gray-700">${content}</p>
                        <div class="flex justify-between items-center mt-3 text-xs text-gray-500">
                            <span>${currentDate}</span>
                            <span>${category.charAt(0).toUpperCase() + category.slice(1)}</span>
                        </div>
                    `;
                    
                    notesContainer.prepend(newNote);
                    
                    // Reset form and close modal
                    document.getElementById('note-title').value = '';
                    document.getElementById('note-content').value = '';
                    addNoteModal.classList.add('hidden');
                }
            });
        
            //refresh
            window.onload = function () {
                window.scrollTo({ top: 0, behavior: 'smooth' });
            };

        // Initialize progress on page load
        updateProgress();