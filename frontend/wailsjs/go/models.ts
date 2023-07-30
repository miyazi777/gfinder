export namespace main {
	
	export class Resource {
	    name: string;
	    info: string;
	    target: string;
	    tag: string[];
	
	    static createFrom(source: any = {}) {
	        return new Resource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.info = source["info"];
	        this.target = source["target"];
	        this.tag = source["tag"];
	    }
	}

}

