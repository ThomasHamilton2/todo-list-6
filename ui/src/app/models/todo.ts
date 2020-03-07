export class Todo {
    id: number;
    title: string = '';
    complete: boolean = false;

    //todo thomas - may delete this? and remove test
    constructor(values: Object = {}) {
      Object.assign(this, values);
    }
}
