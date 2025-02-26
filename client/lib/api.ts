export const fetcher = (...args: [RequestInfo, RequestInit?]) => fetch(...args).then(res => res.json())

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
