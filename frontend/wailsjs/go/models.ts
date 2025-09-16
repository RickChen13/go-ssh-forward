export namespace bll {
	
	export class Result {
	    result: boolean;
	    msg: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new Result(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.result = source["result"];
	        this.msg = source["msg"];
	        this.data = source["data"];
	    }
	}

}

export namespace forward {
	
	export class SshServerConfig {
	    host: string;
	    user: string;
	    pass: string;
	    key_path: string;
	    pass_phrase: string;
	
	    static createFrom(source: any = {}) {
	        return new SshServerConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.host = source["host"];
	        this.user = source["user"];
	        this.pass = source["pass"];
	        this.key_path = source["key_path"];
	        this.pass_phrase = source["pass_phrase"];
	    }
	}
	export class ForwardConfig {
	    id: number;
	    name: string;
	    remote_addr: string;
	    local_addr: string;
	    tag: string;
	
	    static createFrom(source: any = {}) {
	        return new ForwardConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.remote_addr = source["remote_addr"];
	        this.local_addr = source["local_addr"];
	        this.tag = source["tag"];
	    }
	}
	export class Config {
	    forward: ForwardConfig;
	    ssh_server: SshServerConfig;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.forward = this.convertValues(source["forward"], ForwardConfig);
	        this.ssh_server = this.convertValues(source["ssh_server"], SshServerConfig);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

export namespace forwardRule {
	
	export class ConfigForwardRuleUpdateData {
	    name: string;
	    remote_addr: string;
	    local_addr: string;
	    tag: string;
	    sort: number;
	    css_id: number;
	
	    static createFrom(source: any = {}) {
	        return new ConfigForwardRuleUpdateData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.remote_addr = source["remote_addr"];
	        this.local_addr = source["local_addr"];
	        this.tag = source["tag"];
	        this.sort = source["sort"];
	        this.css_id = source["css_id"];
	    }
	}

}

export namespace sshServer {
	
	export class ConfigSshServerC {
	    name: string;
	    host: string;
	    user: string;
	    pass: string;
	    key_path: string;
	    pass_phrase: string;
	    sort: number;
	
	    static createFrom(source: any = {}) {
	        return new ConfigSshServerC(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.host = source["host"];
	        this.user = source["user"];
	        this.pass = source["pass"];
	        this.key_path = source["key_path"];
	        this.pass_phrase = source["pass_phrase"];
	        this.sort = source["sort"];
	    }
	}
	export class ConfigSshServerU {
	    name: string;
	    host: string;
	    user: string;
	    pass: string;
	    key_path: string;
	    pass_phrase: string;
	    sort: number;
	    clear_pass: boolean;
	    clear_pass_phrase: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ConfigSshServerU(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.host = source["host"];
	        this.user = source["user"];
	        this.pass = source["pass"];
	        this.key_path = source["key_path"];
	        this.pass_phrase = source["pass_phrase"];
	        this.sort = source["sort"];
	        this.clear_pass = source["clear_pass"];
	        this.clear_pass_phrase = source["clear_pass_phrase"];
	    }
	}

}

