export namespace models {
	
	export class PasswordEntry {
	    id: string;
	    platform: string;
	    email: string;
	    username: string;
	    description: string;
	    password: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new PasswordEntry(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.platform = source["platform"];
	        this.email = source["email"];
	        this.username = source["username"];
	        this.description = source["description"];
	        this.password = source["password"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
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
	export class PasswordGenerationOptions {
	    length: number;
	    includeUppercase: boolean;
	    includeLowercase: boolean;
	    includeNumbers: boolean;
	    includeSymbols: boolean;
	    useFingerprint: boolean;
	
	    static createFrom(source: any = {}) {
	        return new PasswordGenerationOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.length = source["length"];
	        this.includeUppercase = source["includeUppercase"];
	        this.includeLowercase = source["includeLowercase"];
	        this.includeNumbers = source["includeNumbers"];
	        this.includeSymbols = source["includeSymbols"];
	        this.useFingerprint = source["useFingerprint"];
	    }
	}

}

