class GetTicket {
    constructor(id, { price_in_cents, title, description, discount }) {
        this.id = id;
        this.price_in_cents = price_in_cents;
        this.title = title;
        this.description = description;
        const [pct, value_in_cents] = normalizeTuple(discount)
        this.discount = {
            pct,
            value_in_cents,
        }
    }
}

module.exports = {
    GetTicket,
}