import React, {ChangeEvent, useState} from 'react'
import Head from 'next/head'
import FormControl, {formControlClasses} from '@mui/material/FormControl';
import TextField, {textFieldClasses} from '@mui/material/TextField';
import { styled } from '@mui/material/styles';
import Link from 'next/link';
import { useRouter } from "next/router";
import CircularProgress from '@mui/material/CircularProgress';
import { BASE_URL, TOKEN } from '../utils/constants';
import Swal from 'sweetalert2';

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

export const Signup = () => {
  const router = useRouter();

  const [loading, setLoading] = useState(false);
  const [disabled, setDisabled] = useState(false);
  
  const [form, setForm] = useState({
    first_name: '',
    last_name: '',
    email: '',
    password: '',
    roles: ["INITIATOR"],
  });

  const updateForm = (value: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>, field: string) => {
    let _form: any = {...form}
    _form[field] = value.target.value
    setForm(_form)
  }
  
  const signup = async () => {
    console.log({form})
    setLoading(true)
    setDisabled(true)
    try {
      const { first_name, last_name, email, password, roles } = form;
      const url = `${BASE_URL}/auth/signup`;
      const requestInit = {
        method: 'POST',
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          first_name,
          last_name,
          email,
          password,
          roles
        })
      }
      const response = await fetch(url, requestInit);
      const user = await response.json();
      login()
    } catch (error) {
      console.log({error})
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: `${error}`,
      });
      setLoading(false)
      setDisabled(false)
    }
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
      setLoading(false)
      setDisabled(false)
      router.push(`/${query?.to ? query.to : ''}`);
    } catch (error) {
      console.log({error})
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: `${error}`,
      });
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
          <h1 className='text-dark_blue-100 text-4xl text-center font-semibold tracking-[-1px]'>Create an account</h1>
        </div>

        <div className='w-full p-6 mt-6 mx-6'>
          <StyledFormControl>
            <StyledTextField 
              id="outlined-basic-first_name" 
              label="First name"
              variant="outlined"
              type='text'
              onChange={(e) => updateForm(e, 'first_name')}
              disabled={disabled}
            />
          </StyledFormControl>

          <StyledFormControl>
            <StyledTextField 
              id="outlined-basic-last_name" 
              label="Last name"
              variant="outlined"
              type='text'
              onChange={(e) => updateForm(e, 'last_name')}
              disabled={disabled}
            />
          </StyledFormControl>

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
            className='flex justify-center items-center bg-blue-100 w-full h-14 rounded-2xl text-white font-medium leading-6 tracking-[-0.3px] disabled:opacity-50'
            disabled={disabled}
            onClick={() => signup()}
          >
            <span className='mr-4'>Sign up</span>
            {loading && (
              <CircularProgress size={20} color='inherit' />
            )}
          </button>

          <div
            className='w-full h-14 flex justify-center items-center rounded-2xl text-dark_blue-100 leading-6 tracking-[-0.3px] disabled:opacity-50'
          >
            <span className='mr-1'>Have an account?</span>
            <Link href={`/login${router.query?.to ? `?to=${router.query.to}` : ''}`}>
              <button
                className='text-blue-100 font-medium '
                disabled={disabled}
              >Login</button>
            </Link>
          </div>
        </div>
      </div>
    </div>
  )
}


export default Signup;