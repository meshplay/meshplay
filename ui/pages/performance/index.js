import React, { useEffect } from 'react';
import { NoSsr } from '@material-ui/core';
import MeshplayPerformanceComponent from '../../components/MeshplayPerformance/Dashboard';
import { updatepagepath } from '../../lib/store';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import Head from 'next/head';
import { getPath } from '../../lib/path';

function Performance({ updatepagepath }) {
  useEffect(() => {
    updatepagepath({ path: getPath() });
  }, [updatepagepath]);

  return (
    <NoSsr>
      <Head>
        <title>Performance Dashboard | Meshplay</title>
      </Head>
      <MeshplayPerformanceComponent />
    </NoSsr>
  );
}

const mapDispatchToProps = (dispatch) => ({
  updatepagepath: bindActionCreators(updatepagepath, dispatch),
});

export default connect(null, mapDispatchToProps)(Performance);
