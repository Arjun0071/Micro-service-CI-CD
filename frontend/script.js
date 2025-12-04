// Unified API endpoint via Traefik Ingress (NO PORTS)
const BASE_URL = "http://98.70.14.144";

let TOKEN = localStorage.getItem("token") || null;
let USER_ID = null;

// SECTION SWITCHING
function showSection(id) {
    document.querySelectorAll('.section, .active-section')
        .forEach(sec => sec.classList.remove('active-section'));

    document.getElementById(id).classList.add('active-section');

    // Update tab highlight
    document.querySelectorAll('.nav-link').forEach(link => link.classList.remove('active'));
    document.querySelector(`a[onclick="showSection('${id}')"]`).classList.add('active');
}

// LOGIN
async function login() {
    const email = document.getElementById("login_email").value;
    const password = document.getElementById("login_password").value;

    const res = await fetch(`${BASE_URL}/login`, {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({email, password})
    });

    const data = await res.json();

    if (data.token) {
        TOKEN = data.token;
        localStorage.setItem("token", TOKEN);

        const payload = JSON.parse(atob(TOKEN.split(".")[1]));
        USER_ID = payload.user_id;

        document.getElementById("login_result").innerHTML =
            `<div class="alert alert-success">Login successful!</div>`;

        // Auto-switch to books section
        showSection("books");
    } else {
        document.getElementById("login_result").innerHTML =
            `<div class="alert alert-danger">${data.error}</div>`;
    }
}

// REGISTER USER
async function registerUser() {
    const username = document.getElementById("reg_username").value;
    const email = document.getElementById("reg_email").value;
    const password = document.getElementById("reg_password").value;

    const res = await fetch(`${BASE_URL}/users`, {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({username, email, password})
    });

    const data = await res.json();

    document.getElementById("register_result").innerHTML =
        `<div class="alert alert-info">${data.message || JSON.stringify(data)}</div>`;
}

// LOAD BOOKS
async function loadBooks() {
    const res = await fetch(`${BASE_URL}/books`);
    const books = await res.json();

    const list = document.getElementById("books_list");
    list.innerHTML = "";

    if (books.length === 0) {
        list.innerHTML = `<li class="list-group-item">No books available.</li>`;
        return;
    }

    books.forEach(b => {
        const item = document.createElement("li");
        item.className = "list-group-item";
        item.innerHTML =
            `<strong>[ID: ${b.id}] ${b.title}</strong> by ${b.author} — ₹${b.price} (Stock: ${b.stock})`;
        list.appendChild(item);
    });
}

// CREATE ORDER
async function createOrder() {
    const book_id = Number(document.getElementById("order_book_id").value);
    const quantity = Number(document.getElementById("order_quantity").value);

    const res = await fetch(`${BASE_URL}/orders`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + TOKEN
        },
        body: JSON.stringify({book_id, quantity})
    });

    const data = await res.json();
    document.getElementById("order_result").innerHTML =
        `<div class="alert alert-success">
        Order placed successfully!<br>
        Order ID: ${data.id}<br>
        Book ID: ${data.book_id}<br>
        Quantity: ${data.quantity}<br>
        Total Price: ₹${data.total_price}
    </div>`;
}

// LOAD MY ORDERS
async function loadMyOrders() {
    const res = await fetch(`${BASE_URL}/orders/user/${USER_ID}`, {
        headers: {"Authorization": "Bearer " + TOKEN}
    });

    const orders = await res.json();
    const list = document.getElementById("orders_list");
    list.innerHTML = "";

    if (!orders.length) {
        list.innerHTML = `<li class="list-group-item">No orders placed yet.</li>`;
        return;
    }

    orders.forEach(o => {
        const item = document.createElement("li");
        item.className = "list-group-item";
        item.innerHTML =
            `Order #${o.id} — Book: ${o.book_id} | Qty: ${o.quantity} | Total: ₹${o.total_price} | Status: ${o.status}`;
        list.appendChild(item);
    });
}

