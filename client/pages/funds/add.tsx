import React, { ChangeEvent, use, useState, useEffect } from 'react'
import Head from 'next/head'
import LinearProgress, { linearProgressClasses } from '@mui/material/LinearProgress';
import CopyAllIcon from '@mui/icons-material/CopyAllRounded';
import FormControl, {formControlClasses} from '@mui/material/FormControl';
import TextField, {textFieldClasses} from '@mui/material/TextField';
import { styled } from '@mui/material/styles';
import { useQuery, gql, useMutation } from "@apollo/client";
import { useRouter } from "next/router";
import CircularProgress from '@mui/material/CircularProgress';
import Swal from 'sweetalert2';
import { useAuth } from '../../utils/hooks';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';

const StyledLinearProgress = styled(LinearProgress)(({ theme }) => ({
  [`&.${linearProgressClasses.root}`]: {
    width: '50%',
    height: '0.5rem',
    borderRadius: 50,
    backgroundColor: '#4C51C61A',
  },
}));

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

const Mutation = gql`
  mutation startFund($reason: String!, $description: String!) {
    startFund( input: { reason: $reason, description: $description } ) {
      id
      reason
      description
    }
  }
`

export default function Add() {
  const router = useRouter();
  const {token} = useAuth({reroute: true, from: router.asPath});
  
  const [form, setForm] = useState({
    name: '',
    description: '',
    disabled: false
  });
  const [link, setLink] = useState('');
  const [progress, setProgress] = useState(20);

  const [startFund, { data, loading, error}] = useMutation(Mutation);

  const create = () => {
    let _form: any = {...form}
    _form.disabled = true;
    setForm(_form);
    setProgress(85);
    startFund({
      variables: {
        reason: form.name,
        description: form.description
      },
    })
  }

  useEffect(() => {
    if(data) {
      const {id} = data.startFund;
      const copyLink = `${window.origin}/funds/${id}/fund`
      setLink(copyLink);
    }
  }, [data]);
  
  const copyToClipboard = () => {
    if (link.length > 0) {
      navigator.clipboard.writeText(link)
    }
  }

  const updateForm = (value: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>, field: string) => {
    if (field === 'name') setProgress(35)
    if (field === 'description') setProgress(50)
    let _form: any = {...form}
    _form[field] = value.target.value
    setForm(_form)
  }

  if (error) return Swal.fire({
    icon: 'error',
    title: 'Error',
    text: `${error.message}`,
  })

  return (
    <div className="h-screen min-h-screen flex bg-white">
      <Head>
        <title>Fundbox app</title>
        <meta name="description" content="Generated by Fundbox app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className='w-full flex flex-col'>
        <div className='w-full mt-10 mb-6 flex items-center justify-center'>
          <div className='ml-6'>
            <button onClick={() => router.back()} className='flex items-center justify-center bg-blue-10 h-12 w-12 rounded-full'>
              <ArrowBackIcon className='text-blue-100 w-7 h-7' />
            </button>
          </div>
          <div className='flex-1 flex justify-center' style={{marginLeft: '-64px'}}>
            <StyledLinearProgress
              variant="determinate"
              value={progress}
            />
          </div>
        </div>

        <div className='flex flex-1 flex-col items-center'>
          <div className=''>
            <h1 className='text-dark_blue-100 text-4xl text-center font-semibold tracking-[-1px]'>Welcome, xurror</h1>
            <p className='text-dark_blue-80 text-lg text-center'>Enter required fund details below.</p>
          </div>

          <div className='w-full p-6 mt-6 mx-6'>
            <StyledFormControl>
              <StyledTextField 
                id="outlined-basic" 
                label="Name"
                variant="outlined"
                type='text'
                onChange={(e) => updateForm(e, 'name')}
                disabled={form.disabled}
              />
            </StyledFormControl>

            <StyledFormControl>
              <StyledTextField 
                id="outlined-basic" 
                label="Description"
                variant="outlined"
                type='text'
                onChange={(e) => updateForm(e, 'description')}
                disabled={form.disabled}
              />
            </StyledFormControl>

            {link && (
              <div className='bg-light-100 h-28 rounded-3xl mt-5 p-4 flex items-center'>
                <div className='mr-2 flex-1'>
                  <h4 className='text-dark_blue-100 font-medium text-[15px] break-all'>Copy and share link with contributors!</h4>
                  <p className='text-grey_sub_text break-all text-[15px]'>{link}</p>
                </div>
                <button onClick={() => copyToClipboard()} className=''>
                  <CopyAllIcon className='text-dark_blue-100' />
                </button>
              </div>
            )}
          </div>
          <div className='w-full px-6 my-5 mt-20'>
            <button
              className='bg-blue-100 w-full h-14 rounded-2xl text-white font-medium leading-6 tracking-[-0.3px] disabled:opacity-50'
              disabled={form.disabled || loading}
              onClick={() => create()}
            >
              <span className='mr-4'>Create fund</span>
              {loading && (
                <CircularProgress size={20} color='inherit' />
              )}
            </button>
          </div>
        </div>
      </main>
    </div>
  )
}
