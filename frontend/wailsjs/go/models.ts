export namespace define {
	
	export class Connection {
	    identity: string;
	    name: string;
	    addr: string;
	    port: string;
	    username: string;
	    password: string;
	    type: string;
	    ssh_addr: string;
	    ssh_port: string;
	    ssh_username: string;
	    ssh_password: string;
	
	    static createFrom(source: any = {}) {
	        return new Connection(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.identity = source["identity"];
	        this.name = source["name"];
	        this.addr = source["addr"];
	        this.port = source["port"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.type = source["type"];
	        this.ssh_addr = source["ssh_addr"];
	        this.ssh_port = source["ssh_port"];
	        this.ssh_username = source["ssh_username"];
	        this.ssh_password = source["ssh_password"];
	    }
	}
	export class KeyListRequest {
	    conn_identity: string;
	    db: number;
	    keyword: string;
	
	    static createFrom(source: any = {}) {
	        return new KeyListRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.conn_identity = source["conn_identity"];
	        this.db = source["db"];
	        this.keyword = source["keyword"];
	    }
	}
	export class KeyValueRequest {
	    conn_identity: string;
	    db: number;
	    key: string;
	
	    static createFrom(source: any = {}) {
	        return new KeyValueRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.conn_identity = source["conn_identity"];
	        this.db = source["db"];
	        this.key = source["key"];
	    }
	}

}

