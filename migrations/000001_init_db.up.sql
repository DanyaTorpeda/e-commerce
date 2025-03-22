create table if not exists Users (
    id serial primary key,
    first_name varchar(50),
    last_name varchar(50),
    password_hash text not null,
    email varchar(100) not null unique,
    role varchar(10) not null default 'user',
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table if not exists Products (
    id serial primary key,
    name varchar(50) not null,
    description text,
    price decimal not null,
    stock integer,
    created_at timestamp default now(),
    updated_at timestamp default now()
);

create table if not exists Orders (
    id serial primary key,
    user_id integer not null,
    total_price decimal,
    status varchar(20) not null default 'pending',
    payment_status varchar(20) not null default 'pending',
    created_at timestamp default now(),
    updated_at timestamp default now(),
    foreign key (user_id) references Users(id) on delete cascade
);

create table if not exists OrderItems (
    id serial primary key,
    order_id integer not null,
    product_id integer not null,
    quantity integer,
    price decimal,
    foreign key (order_id) references Orders(id) on delete cascade,
    foreign key (product_id) references Products(id) on delete cascade
);

create table if not exists ShoppingCarts (
    id serial primary key,
    user_id integer not null,
    total_price decimal,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    foreign key (user_id) references Users(id) on delete cascade
);

create table if not exists CartItems (
    id serial primary key,
    cart_id integer not null,
    product_id integer not null,
    quantity integer,
    price decimal,
    foreign key (cart_id) references ShoppingCarts(id) on delete cascade,
    foreign key (product_id) references Products(id) on delete cascade
);

create table if not exists Payments (
    id serial primary key,
    order_id integer not null,
    amount integer,
    status varchar(10) not null default 'pending',
    payment_method varchar(10) not null default 'stripe',
    created_at timestamp default now(),
    updated_at timestamp default now(),
    foreign key (order_id) references Orders(id) on delete cascade
);