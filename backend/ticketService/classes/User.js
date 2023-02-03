class User {
    constructor(id, { first_name, last_name, date_of_birth }) {
        this.id = id;
        this.first_name = first_name;
        this.last_name = last_name;
        this.date_of_birth = date_of_birth;
    }

    get fullname() {
        return this.returnFullName();
    }

    returnFullName() {
        return `${this.first_name} ${this.last_name}`;
    }
}

module.exports = User;
