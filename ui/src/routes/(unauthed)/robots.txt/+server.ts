import { text } from '@sveltejs/kit';
import { getOrigin } from '$lib/util';

export const prerender = true;

export async function GET() {
	return text(`User-agent: *
Allow: /

Sitemap: ${getOrigin()}/sitemap.xml`);
}
