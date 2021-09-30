import React from 'react';
import clsx from 'clsx';
import styles from './PageFooter.module.css';

export default function PageFooter() {
  return (
    <footer className={styles.footer}>
      <div className="container">
        <p className="copyright">Copyright © 2021 Wawandco Inc.</p>
      </div>
    </footer>
  );
}
