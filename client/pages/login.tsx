import React, {ChangeEvent, useState} from 'react'
import Head from 'next/head'
import FormControl, {formControlClasses} from '@mui/material/FormControl';
import TextField, {textFieldClasses} from '@mui/material/TextField';
import { styled } from '@mui/material/styles';
import Link from 'next/link';
import { useRouter } from "next/router";
import CircularProgress from '@mui/material/CircularProgress';
import { BASE_URL, TOKEN } from '../utils/constants';

const StyledFormControl = styled(FormControl)(({ theme }) => ({
  [`&.${formControlClasses.root}`]: {
    width: '100%',
    marginBottom: '1rem',
    borderRadius: '12px',
  },
}));

const StyledTextField = styled(TextField)(({ theme }) => ({
  [`&.${textFieldClasses.root}`]: {
    borderRadius: '12px',
  },
}));

export const login = () => {
  const router = useRouter();

  const [loading, setLoading] = useState(false);
  const [disabled, setDisabled] = useState(false);
  const [form, setForm] = useState({
    email: '',
    password: ''
  });

  const updateForm = (value: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>, field: string) => {
    let _form: any = {...form}
    _form[field] = value.target.value
    setForm(_form)
  }
  
  const login = async () => {
    const query = router.query;

    try {
      const { password, email } = form;
      const url = `${BASE_URL}/auth/login`;
      const requestInit = {
        method: 'POST',
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          email,
          password,
        })
      }
      const response = await fetch(url, requestInit);
      const data = await response.json();

      localStorage.setItem(TOKEN, data.token);
      router.push(`/${query?.to ? query.to : ''}`);
    } catch (error) {
      console.log({error})
      window.alert(error)
      setLoading(false)
      setDisabled(false)
    }
  }

  return (
    <div className="h-screen min-h-screen flex bg-white">
      <Head>
        <title>Fundbox app</title>
        <meta name="description" content="Generated by Fundbox app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <div className='flex flex-1 flex-col items-center justify-center'>
        <div className='mt-10'>
          <h1 className='text-dark_blue-100 text-4xl text-center font-semibold tracking-[-1px]'>Login</h1>
        </div>

        <div className='w-full p-6 mt-6 mx-6'>
          <StyledFormControl>
            <StyledTextField 
              id="outlined-basic-email" 
              label="Email"
              variant="outlined"
              type='email'
              onChange={(e) => updateForm(e, 'email')}
              disabled={disabled}
            />
          </StyledFormControl>

          <StyledFormControl>
            <StyledTextField 
              id="outlined-basic-password" 
              label="Password"
              variant="outlined"
              type='password'
              onChange={(e) => updateForm(e, 'password')}
              disabled={disabled}
            />
          </StyledFormControl>
        </div>

        <div className='w-full px-6 my-5 mt-5'>
          <button
            className='bg-blue-100 w-full h-14 rounded-2xl text-white font-medium leading-6 tracking-[-0.3px] disabled:opacity-50'
            disabled={disabled}
            onClick={() => login()}
          >
            <span className='mr-4'>Login</span>
            {loading && (
              <CircularProgress size={20} color='inherit' />
            )}
          </button>

          <div
            className='w-full h-14 flex justify-center items-center rounded-2xl text-dark_blue-100 leading-6 tracking-[-0.3px] disabled:opacity-50'
          >
            <span className='mr-1'>Done't an account?</span>
            <Link href={`/signup${router.query?.to ? `?to=${router.query.to}` : ''}`}>
              <button
                className='text-blue-100 font-medium '
                disabled={disabled}
              >Sign up</button>
            </Link>
          </div>
        </div>
      </div>
    </div>
  )
}


export default login;