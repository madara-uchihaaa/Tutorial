import { v4 as uuidv4 } from 'uuid';
export class Chat {
    id!: number;
    userId: string = "";
    question!: string;
    answer!: string;
    image!: string;

    constructor(json: any) {
        if (json) {
            Object.assign(this, {
                ...json,
                id: json.id ? json.id : uuidv4(),
            });
        }
    }
}