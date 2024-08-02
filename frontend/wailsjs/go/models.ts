export namespace systemModel {
	
	export class Setting {
	    id: number;
	    authOpened: boolean;
	    isShow: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Setting(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.authOpened = source["authOpened"];
	        this.isShow = source["isShow"];
	    }
	}

}

export namespace todoModel {
	
	export class Todo {
	    id: number;
	    title: string;
	    content: string;
	    create_at: number;
	    completed: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Todo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.content = source["content"];
	        this.create_at = source["create_at"];
	        this.completed = source["completed"];
	    }
	}

}

