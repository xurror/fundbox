/**
 * Custom fetcher for SWR that properly handles HTTP error responses
 * @param url - The URL to fetch
 * @param init - Optional fetch init options
 * @returns The parsed JSON response
 * @throws FetchError with status, info, and message when the server returns an error
 */
export class FetchError<T> extends Error {
  status: number;
  info: T;
  
  constructor(message: string, status: number, info: T) {
    super(message);
    this.name = 'FetchError';
    this.status = status;
    this.info = info;
  }
}

export const fetcher = async <T>(
  url: string, 
  init?: RequestInit
): Promise<T> => {
  const response = await fetch(url, {
    ...init,
    headers: {
      'Content-Type': 'application/json',
      ...init?.headers,
    },
  });

  // First, check if the response is OK (status in the range 200-299)
  if (!response.ok) {
    // Try to parse error information from the response
    let errorInfo;
    try {
      errorInfo = await response.json();
    } catch (e) {
      console.error('Failed to parse error response:', e);
      // If parsing fails, use a simple error message
      errorInfo = { message: 'An error occurred while fetching the data.' };
    }

    // Throw a custom error with status code and parsed error info
    throw new FetchError(
      errorInfo.message || `API error with status code: ${response.status}`,
      response.status,
      errorInfo
    );
  }
  
  // If response is OK, parse and return the JSON data
  return response.json();
};

export const mapQueryParams = (params: { [key: string]: string }) => {
  return Object.keys(params)
    .map((key) => `${key}=${params[key]}`)
    .join("&")
}

export const mutator = async (url: string, { arg }: { arg: string }) => {
  return fetch(url, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: arg
  }).then(res => res.json())
}
