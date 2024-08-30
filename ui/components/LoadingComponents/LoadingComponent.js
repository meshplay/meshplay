import { makeStyles, Typography, useTheme } from '@material-ui/core';
import clsx from 'clsx';
import React from 'react';
import AnimatedMeshSync from './Animations/AnimatedMeshSync';
import AnimatedMeshPattern from './Animations/AnimatedMeshPattern';
import AnimatedMeshplay from './Animations/AnimatedMeshplay';
import AnimatedFilter from './Animations/AnimatedFilter';
import PropTypes from 'prop-types';
import AnimatedLightMeshplay from './Animations/AnimatedLightMeshplay';

const useStyles = makeStyles(() => ({
  loadingWrapper: {
    textAlign: 'center',
    marginTop: 'calc(50vh - 141px)',
    transform: 'translateY(-50%)',
  },
}));

const animatedIconList = {
  AnimatedMeshPattern: <AnimatedMeshPattern style={{ height: '100px', margin: '4px 0px 8px' }} />,
  AnimatedMeshSync: (
    <>
      <AnimatedMeshSync style={{ height: '75px' }} />
      <img
        src="/static/img/meshplay-logo/meshplay-black.svg"
        alt="mehsery-logo"
        width="125px"
        style={{ margin: '4px 0px 8px' }}
      />
    </>
  ),
  AnimatedFilter: <AnimatedFilter style={{ height: '75px', margin: '4px 0px 8px' }} />,
  AnimatedMeshplay: <AnimatedMeshplay style={{ height: '100px', margin: '4px 0px 8px' }} />,
};

const animatedLightIconList = {
  AnimatedMeshPattern: <AnimatedMeshPattern style={{ height: '100px', margin: '4px 0px 8px' }} />,
  AnimatedMeshSync: (
    <>
      <AnimatedMeshSync style={{ height: '75px' }} />
      <img
        src="/static/img/meshplay-logo/meshplay-white.svg"
        alt="mehsery-logo"
        width="125px"
        style={{ margin: '4px 0px 8px' }}
      />
    </>
  ),
  AnimatedFilter: <AnimatedFilter style={{ height: '75px', margin: '4px 0px 8px' }} />,
  AnimatedMeshplay: <AnimatedLightMeshplay style={{ height: '100px', margin: '4px 0px 8px' }} />,
};
function LoadingScreen(props) {
  const { message, className, animatedIcon, ...other } = props;
  const classes = useStyles();
  const theme = useTheme();
  return (
    <div className={clsx(classes.loadingWrapper, className)} {...other}>
      {theme.palette.type === 'light'
        ? animatedIconList[animatedIcon]
        : animatedLightIconList[animatedIcon]}
      <Typography variant="caption" component="div">
        {message}
      </Typography>
    </div>
  );
}

LoadingScreen.propTypes = {
  message: PropTypes.string.isRequired,
  className: PropTypes.string,
  animatedIcon: PropTypes.string.isRequired,
};

export default LoadingScreen;
