import { MAX_INT, random } from '@aicacia/rand';

export function createInsecureID() {
	return (random() * MAX_INT) | 0;
}

export function getSearchTerms(search: string): string[] {
	return search
		.toLowerCase()
		.trim()
		.split(/\s+/)
		.filter((t) => !!t);
}

export function toURLSafe(value: string): string {
	return value
		.trim()
		.toLowerCase()
		.replace(/[\s]+/gi, '-')
		.replace(/[^\w\d\-_]+/gi, '');
}

export function getId<T extends { id: number }>(value: T) {
	return value.id;
}
export function getName<T extends { name: string }>(value: T) {
	return value.name;
}
