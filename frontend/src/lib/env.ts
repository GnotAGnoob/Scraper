type Env = {
  apiBaseUrl: string
}

const apiBaseUrl = import.meta.env.VITE_API_BASE_URL

if (!apiBaseUrl) {
  throw new Error(
    'VITE_API_BASE_URL is not defined. Add it to your .env as VITE_API_BASE_URL.',
  )
}

// Ensure protocol is present
const normalizedApiBaseUrl = apiBaseUrl.startsWith('http://') || apiBaseUrl.startsWith('https://')
  ? apiBaseUrl
  : `http://${apiBaseUrl}`

export const env: Env = {
  apiBaseUrl: normalizedApiBaseUrl,
}



