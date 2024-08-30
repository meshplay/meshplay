import React, { useEffect } from 'react';
import { NoSsr, Paper, withStyles } from '@material-ui/core';
import MeshplayResults from '../../components/MeshplayResults';
import { updatepagepath } from '../../lib/store';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';
import Head from 'next/head';
import { getPath } from '../../lib/path';

const styles = { paper: { maxWidth: '90%', margin: 'auto', overflow: 'hidden' } };

function Results({ classes, updatepagepath }) {
  useEffect(() => {
    updatepagepath({ path: getPath() });
  }, [updatepagepath]);

  return (
    <NoSsr>
      <Head>
        <title>Performance Test Results | Meshplay</title>
      </Head>
      <Paper className={classes.paper}>
        <MeshplayResults />
      </Paper>
    </NoSsr>
  );
}

const mapDispatchToProps = (dispatch) => ({
  updatepagepath: bindActionCreators(updatepagepath, dispatch),
});

export default withStyles(styles)(connect(null, mapDispatchToProps)(Results));
