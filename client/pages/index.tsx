import Head from 'next/head'
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import { useState } from 'react';
import LinearProgress, { linearProgressClasses } from '@mui/material/LinearProgress';

import AddIcon from '@mui/icons-material/Add';
import PersonIcon from '@mui/icons-material/PersonOutlined';
import LoginIcon from '@mui/icons-material/Login';
import LogoutIcon from '@mui/icons-material/Logout';
import ViewIcon from '@mui/icons-material/FolderOpenOutlined';
import Link from 'next/link';
import Popover, {popoverClasses} from '@mui/material/Popover';
import { styled } from '@mui/material/styles';
import { useAuth } from '../utils/hooks';

const StyledPopover = styled(Popover)(({ theme }) => ({
  [`&.${popoverClasses.paper}`]: {
    backgroundColor: 'red',
    borderRadius: 14,
  },
}));

const StyledLinearProgress = styled(LinearProgress)(({ theme }) => ({
  [`&.${linearProgressClasses.root}`]: {
    width: '50%',
    height: '0.5rem',
    borderRadius: 50,
    backgroundColor: '#4C51C61A',
  },
  [`&.${linearProgressClasses.bar1Determinate}`]: {
    backgroundColor: 'red',
  },
}));

export default function App() {
  const [token] = useAuth({reroute: false});
  const [auth, setAuth] = useState(false)
  const [visible, setVisible] = useState(false)
  const [anchorEl, setAnchorEl] = useState<HTMLButtonElement | null>(null);

  const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
    console.log("clicked")
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const open = Boolean(anchorEl);
  const id = open ? 'simple-popover' : undefined;

  console.log({token})
  return (
    <div className="h-screen min-h-screen flex bg-white">
      <Head>
        <title>Fundbox app</title>
        <meta name="description" content="Generated by Fundbox app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className='w-full flex flex-col'>
        <div className='w-full mt-10 mb-5 flex items-center justify-center'>
          <StyledLinearProgress
            variant="determinate"
            value={10}
            style={token && token.length > 0 ? {marginRight: '-64px', marginLeft: 'auto'} : {}}
          />
          {token && token.length > 0 && (
            <div className='ml-auto mr-6'>
              <button aria-describedby={id} onClick={handleClick} className='flex items-center justify-center bg-blue-10 h-10 w-10 rounded-full'>
                <PersonIcon className='text-blue-100 w-7 h-7' />
              </button>
              <StyledPopover
                id={id}
                open={open}
                anchorEl={anchorEl}
                onClose={handleClose}
                anchorOrigin={{
                  vertical: 'bottom',
                  horizontal: 'center',
                }}
                transformOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
              >
                <div className='py-2 w-48'>
                  <button className='w-full text-start px-4 text-blue-100 font-semibold h-10 flex items-center border-b border-dark_blue-20'>
                    <Link href='/funds'>
                      <ViewIcon className='text-blue-100 w-5 h-5 mr-3' /><span>View Funds</span>
                    </Link>
                  </button>
                  <button className='w-full text-start px-4 text-blue-100 font-semibold h-10 flex items-center'>
                    <LogoutIcon className='text-blue-100 w-5 h-5 mr-3' /><span>Log out</span>
                  </button>
                </div>
              </StyledPopover>
            </div>
          )}
        </div>

        <div className='flex flex-1 flex-col items-center'>
          <h1 className='text-dark_blue-100 text-3xl font-semibold tracking-[-1px]'>Create a Fund</h1>

          <div className='flex flex-col flex-1 items-center p-6 mt-6'>
            {!token && (
              <div className='bg-blue-100 p-6 rounded-[40px] h-56 flex flex-col items-center mb-5 w-full'>
                <h3 className='text-white text-2xl font-semibold text-center leading-8 mb-1'>Already have a fund.</h3>
                <p className='text-white-80 text-center leading-6 text-sm tracking-[-0.3px] px-1'>Login to track fund progress, see contributors and other information.</p>

                <div className='flex flex-1 justify-center items-end'>
                  <Link href='/login'>
                    <button className='flex justify-center items-center'>
                        <div className='flex items-center justify-center bg-white-10 h-14 w-14 rounded-full mr-2'>
                          <LoginIcon className='text-white w-6 h-6' />
                        </div>
                        <span className='text-white font-medium leading-7 tracking-[-0.4px]'>Login</span>
                    </button>
                  </Link>
                </div>
              </div>
            )}

            <div className='bg-light-100 p-6 rounded-[40px] h-56 flex flex-col items-center'>
              <h3 className='text-dark_blue-100 text-2xl font-semibold text-center leading-8'>Create a Fund</h3>
              <p className='text-dark_blue-80 text-center leading-6' >Simply create a fund, copy and share the link with your friends to start collecting!</p>

              <div className='flex flex-1 justify-center items-end'>
                <Link href='/funds/add'>
                  <button className='flex justify-center items-center focus:outline-none focus:shadow-none'>
                    <div className='flex items-center justify-center bg-blue-10 h-14 w-14 rounded-full mr-2'>
                      <AddIcon className='text-blue-100 w-7 h-7' />
                    </div>
                    <span className='text-blue-100 font-medium leading-7 tracking-[-0.4px]'>Add</span>
                  </button>
                </Link>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  )
}
