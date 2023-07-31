export namespace main {
	
	export class InnerResource {
	    name: string;
	    info: string;
	    target: string;
	    tag: string;
	    command: string;
	    command_args: string[];
	
	    static createFrom(source: any = {}) {
	        return new InnerResource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.info = source["info"];
	        this.target = source["target"];
	        this.tag = source["tag"];
	        this.command = source["command"];
	        this.command_args = source["command_args"];
	    }
	}

}

