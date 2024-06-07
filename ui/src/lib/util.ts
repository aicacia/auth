import { MAX_INT, random, fillBytes } from '@aicacia/rand';

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

export function formatPhoneNumber(phoneNumberString: string) {
	const cleaned = ('' + phoneNumberString).replace(/\D/g, '');
	const match = cleaned.match(/^(\d{1}|\d{2}|\d{3})?(\d{3})(\d{3})(\d{4})$/);
	if (match) {
		const intlCode = match[1] ? `+${match[1]} ` : '';
		return `${intlCode}(${match[2]}) ${match[3]}-${match[4]}`;
	}
	return phoneNumberString;
}

export function generateRandomBase64String(size: number): string {
	const bytes = fillBytes(new Uint8Array(size));
	let buffer = new Array(size);
	for (let i = 0; i < bytes.length; i++) {
		buffer.push(String.fromCharCode(bytes[i]));
	}
	return btoa(buffer.join(''));
}
