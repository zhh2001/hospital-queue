const queues = document.getElementById('queues');
const emptyMessage = document.getElementById('emptyMessage');

const name = document.getElementById('name');
const phone = document.getElementById('phone');
const department = document.getElementById('department');

const newQueueButton = document.getElementById('newQueueButton');
newQueueButton.addEventListener('click', createNewQueue);

loadData();
setInterval(loadData, 5000);

function loadData() {
    fetch('/queue/', {
                method: 'GET',
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('获取队列数据失败');
                }
                return response.json();
            })
            .then(data => {
                // 先清空现有内容
                queues.innerHTML = '';

                // 如果没有等待队列，显示提示信息
                if (data.length === 0) {
                    queueItemsEl.innerHTML = "";
                    emptyMessage.className = "empty-message";
                    return;
                }

                emptyMessage.className = "empty-message hidden";

                for (let i = 0; i < data.length; i++) {
                    const queueItem = document.createElement('tr');
                    queueItem.innerHTML = `
                        <td>
                            <button type="button" class="callBtn" data-id=${data[i].id}>叫号</button>
                        </td>
                        <td class=${data[i].status === 0 ? "" : "gray"}>${data[i].status === 0 ? "未叫号" : "已叫号"}</td>
                        <td>${data[i].number}</td>
                        <td>${data[i].name}</td>
                        <td>${data[i].phone}</td>
                        <td>${data[i].department}</td>
                        <td>${data[i].remaining}</td>
                        <td>${data[i].datetime}</td>
                    `;
                    queues.appendChild(queueItem);
                }

                const buttons = document.getElementsByClassName('callBtn');
                for (let i = 0; i < buttons.length; i++) {
                    const button = buttons[i];
                    button.addEventListener('click', () => {
                         callQueue(button.dataset.id);
                    });
                }
            })
            .catch(error => {
                alert(error.message);
            });
}

function createNewQueue() {
    console.log(name.value, phone.value, department.value)
    if (name.value.trim() === "") {
        alert("患者姓名不能为空！")
        return
    }

    if (department.value <= 0) {
        alert("诊室号必须为正整数！")
        return
    }

    fetch('/queue/new', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: `name=${name.value}&phone=${phone.value}&department=${department.value}`
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('创建排队号码失败');
                }
                return response.json();
            })
            .then(data => {
                name.value = ""
                phone.value = ""
                alert(`成功创建排队号码: ${data.number}`);
                loadData(); // 刷新数据
            })
            .catch(error => {
                alert(error.message);
            });
}

function callQueue(queueID) {
    fetch('/queue/call', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: `id=${queueID}`
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('叫号失败');
                }
                alert(`成功叫号`);
                loadData(); // 刷新数据
            })
            .catch(error => {
                alert(error.message);
            });
}