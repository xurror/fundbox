import { useState, useEffect } from "react";
import { TOKEN } from './constants';
import { useRouter } from "next/router";

export const useAuth = (settings: {reroute: boolean}) => {
  const router = useRouter();
  const [data, setData] = useState<string | null>(null);

  useEffect(() => {
    const token = localStorage.getItem(TOKEN);
    //  TODO: could also check if token is expired
    if (token) setData(token)

    if (settings.reroute && !token) router.push('/login');
  }, []);

  return [data];
};

