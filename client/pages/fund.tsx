import Head from 'next/head'
import { ChangeEvent, useState } from 'react';
import AddIcon from '@mui/icons-material/Add';
import PaymentsIcon from '@mui/icons-material/PaymentsOutlined';
import DoneIcon from '@mui/icons-material/DoneAllOutlined';
import Radio from '@mui/material/Radio';

import Steps from '../components/Steps'


const steps = [
  {
    label: 'Add your name',
    description: ``,
    icon: <AddIcon />,

  },
  {
    label: 'Pay using one of the following methods',
    description: '',
    icon: <PaymentsIcon />
  },
  {
    label: 'Make Payment',
    description: '',
    icon: <PaymentsIcon />
  },
  {
    label: 'Completed',
    description: '',
    icon: <DoneIcon />
  }
];

export const Fund = () => {
  const [currentStep, setCurrentStep] = useState(0);
  const [fundName, setFundName] = useState('');
  const [paymentMethod, setPaymentMethod] = useState('');

  const goNextStep = () => {
    if (currentStep === steps.length - 1) return;
    setCurrentStep(currentStep + 1)
  }

  const goBackStep = () => {
    if (currentStep === 0) return;
    setCurrentStep(currentStep - 1)
  }

  const setName = (e: ChangeEvent<HTMLInputElement>) => {
    setFundName(e.target.value);
  }

  const handlePaymentSelect = (value: any) => {
    setPaymentMethod(value);
  };

  return (
    <div className="h-screen min-h-screen flex flex-col justify-center items-center px-[200px] py-16 bg-primary-light">
      <Head>
        <title>Fundbox - Fund</title>
        <meta name="description" content="Generated by Fundbox app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex flex-col flex-1 bg-basic-200 w-full rounded-3xl shadow-md shadow px-20">
        <div className='py-14 border-b border-basic-500'>
          <h1 className="text-4xl">Fund for <span className='font-bold text-secondary-dark'>xurror</span></h1>
          <p className="text-md text-basic-600 pt-1">Contribute to this fund by completing the form with your information</p>
        </div>

        <div className='flex flex-1'>
          <div className='w-1/3 border-r border-basic-500 px-5 py-10'>
            <Steps steps={steps} currentStep={currentStep} />
          </div>
          <div className='w-2/3 p-10 flex flex-col'>
            <div className='flex-1'>
              {currentStep === 0 && (
                <div>
                  <div className='text-md text-basic-600 font-medium pb-2'>Name</div>
                  <input value={fundName} onChange={setName} className='h-14 px-3 border text-basic-600 w-full rounded-md border-basic-500 bg-basic-200' type="text" />
                </div>
              )}
              {currentStep === 1 && (
                <div>
                  <div className='text-md text-basic-600 font-medium pb-8'>Payment Methods</div>
                  <div className=''>
                    <div className='flex mb-4'>
                      <button className='flex flex-1 items-center border rounded-md shadow shadow-primary-light h-28 pl-3 border-basic-500' onClick={() => handlePaymentSelect('momo')}>MOMO</button>
                      <Radio
                        checked={paymentMethod === 'momo'}
                        onChange={(e) => handlePaymentSelect(e.target.value)}
                        value="momo"
                        name="radio-buttons"
                        inputProps={{ 'aria-label': 'MOMO' }}
                      />
                    </div>
                    <div className='flex mb-4'>
                      <button className='flex flex-1 items-center border rounded-md shadow shadow-primary-light h-28 pl-3 border-basic-500' onClick={() => handlePaymentSelect('om')}>OM</button>
                      <Radio
                        checked={paymentMethod === 'om'}
                        onChange={(e) => handlePaymentSelect(e.target.value)}
                        value="om"
                        name="radio-buttons"
                        inputProps={{ 'aria-label': 'OM' }}
                      />
                    </div>
                  </div>
                </div>
              )}
            </div>

            <div className='flex justify-between items-center'>
              <button onClick={() => goBackStep()} disabled={currentStep === 0} className='border rounded-md h-9 border-basic-500 w-28 hover:border-primary text-basic-600 transition duration-200 delay-75 disabled:cursor-not-allowed disabled:opacity-50 disabled:hover:border-basic-500'>Back</button>
              <button onClick={() => goNextStep()} disabled={currentStep === steps.length-1} className='bg-primary border rounded-md h-9 border-primary w-28 hover:bg-primary-light hover:text-dark-100 text-basic-100 transition duration-200 delay-75 disabled:cursor-not-allowed disabled:opacity-50 disabled:hover:bg-primary disabled:hover:text-basic-100'>Next</button>
            </div>
          </div>
        </div>
      </main>
    </div>
  )
}

export default Fund;