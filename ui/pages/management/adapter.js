import { NoSsr } from '@material-ui/core';
import Head from 'next/head';
import { withRouter } from 'next/router';
import React from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import MeshplayPlayComponent from '../../components/MeshplayPlayComponent';
import { updatepagepath } from '../../lib/store';

const Manage = () => {
  return (
    <NoSsr>
      <Head>
        <title>Management | Meshplay </title>
      </Head>
      <MeshplayPlayComponent />
    </NoSsr>
  );
};

const mapDispatchToProps = (dispatch) => {
  return { updatepagepath: bindActionCreators(updatepagepath, dispatch) };
};

export default connect(null, mapDispatchToProps)(withRouter(Manage));
