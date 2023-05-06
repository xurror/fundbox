import * as React from 'react';
import PropTypes from 'prop-types';
import { styled } from '@mui/material/styles';
import Stack from '@mui/material/Stack';
import Stepper from '@mui/material/Stepper';
import Step from '@mui/material/Step';
import StepLabel from '@mui/material/StepLabel';
import AddIcon from '@mui/icons-material/Add';
import ShareIcon from '@mui/icons-material/Share';
import VideoLabelIcon from '@mui/icons-material/VideoLabel';
import StepConnector, { stepConnectorClasses } from '@mui/material/StepConnector';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import StepContent from '@mui/material/StepContent';
import Box from '@mui/material/Box';
import { FC, useEffect, useState } from 'react';

const ColorlibConnector = styled(StepConnector)(({ theme }) => ({
  [`&.${stepConnectorClasses.alternativeLabel}`]: {
    top: 22,
  },
  [`&.${stepConnectorClasses.active}`]: {
    [`& .${stepConnectorClasses.line}`]: {
      backgroundImage:
        'linear-gradient( 95deg,#C6AAE0 0%, #8B71CD 50%, #412DB3 100%)',
    },
  },
  [`&.${stepConnectorClasses.completed}`]: {
    [`& .${stepConnectorClasses.line}`]: {
      backgroundImage:
        'linear-gradient( 95deg,#C6AAE0 0%, #8B71CD 50%, #412DB3 100%)',
    },
  },
  [`& .${stepConnectorClasses.line}`]: {
    width: 5,
    marginLeft: 10,
    height: 50,
    border: 0,
    backgroundColor: '#F3F2FE',
    borderRadius: 1,
  },
}));

const ColorlibStepIconRoot = styled('div')(({ theme, ownerState }) => ({
  backgroundColor: '#F3F2FE',
  zIndex: 1,
  color: '#1D1D1F',
  width: 50,
  height: 50,
  display: 'flex',
  borderRadius: '50%',
  justifyContent: 'center',
  alignItems: 'center',
  boxShadow: '0 2px 5px 0 rgba(0,0,0,.10)',
  ...(ownerState.active && {
    backgroundImage:
      'linear-gradient( 136deg, #C6AAE0 0%, #8B71CD 50%, #412DB3 100%)',
    boxShadow: '0 4px 10px 0 rgba(0,0,0,.15)',
    color: '#FAFAFA',
  }),
  ...(ownerState.completed && {
    backgroundImage:
      'linear-gradient( 136deg, #C6AAE0 0%, #8B71CD 50%, #412DB3 100%)',
    color: '#FAFAFA',
  }),
}));

function ColorlibStepIcon(props: any) {
  const { active, completed, className } = props;

  const icons = {
    1: <AddIcon />,
    2: <ShareIcon />,
  };

  return (
    <ColorlibStepIconRoot ownerState={{ completed, active }} className={className}>
      {icons[String(props.icon)]}
    </ColorlibStepIconRoot>
  );
}

interface step {
  label: string,
  description: string
}

const Steps:FC<{steps: step[], currentStep: number}>= ({ steps, currentStep}) =>  {
  const [activeStep, setActiveStep] = useState(currentStep)

  useEffect(() => {
    setActiveStep(currentStep)
  }, [currentStep])
  
  return (
    <Stack sx={{ width: '100%' }}>
      <Stepper activeStep={activeStep} orientation="vertical" connector={<ColorlibConnector />}>
        {steps.map((step, index) => (
          <Step key={step.label}>
            <StepLabel StepIconComponent={ColorlibStepIcon} className='p-0'>
              <div className='text-dark-100 text-sm'>{step.label}</div>
            </StepLabel>
            {/* <StepContent>
              <Typography>{step.description}</Typography>
              <Box sx={{ mb: 2 }}>
                <div>
                  <Button
                    variant="contained"
                    // onClick={handleNext}
                    sx={{ mt: 1, mr: 1 }}
                  >
                    {index === steps.length - 1 ? 'Finish' : 'Continue'}
                  </Button>
                  <Button
                    disabled={index === 0}
                    // onClick={handleBack}
                    sx={{ mt: 1, mr: 1 }}
                  >
                    Back
                  </Button>
                </div>
              </Box>
            </StepContent> */}
          </Step>
        ))}
      </Stepper>
    </Stack>
  );
}

export default Steps;