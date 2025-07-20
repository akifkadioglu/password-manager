export namespace controller {
	
	export class BackupMetadata {
	    passwordCount: number;
	    exportDate: string;
	    version: string;
	    encrypted: boolean;
	
	    static createFrom(source: any = {}) {
	        return new BackupMetadata(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.passwordCount = source["passwordCount"];
	        this.exportDate = source["exportDate"];
	        this.version = source["version"];
	        this.encrypted = source["encrypted"];
	    }
	}
	export class BackupOption {
	    id: string;
	    name: string;
	    description: string;
	    icon: string;
	    available: boolean;
	
	    static createFrom(source: any = {}) {
	        return new BackupOption(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.icon = source["icon"];
	        this.available = source["available"];
	    }
	}
	export class EncryptedBackupData {
	    version: string;
	    exportDate: string;
	    appName: string;
	    encrypted: boolean;
	    encryptedPayload: string;
	    encryptionMethod: string;
	
	    static createFrom(source: any = {}) {
	        return new EncryptedBackupData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.version = source["version"];
	        this.exportDate = source["exportDate"];
	        this.appName = source["appName"];
	        this.encrypted = source["encrypted"];
	        this.encryptedPayload = source["encryptedPayload"];
	        this.encryptionMethod = source["encryptionMethod"];
	    }
	}

}

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

