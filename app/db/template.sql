insert into catalogue (product_id, price, stock, updated_at)
values (?, ?, ?, now())
on duplicate key
    update price=
               values(price),
           stock=
               values(stock),
           updated_at=
               values(updated_at);