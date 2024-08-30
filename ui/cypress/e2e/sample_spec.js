describe('Visit Meshplay', () => {
  it('Visits meshplay', () => {
    cy.selectProviderNone();
    cy.visit('/');
  });
});

describe('Visit Meshplay settings', () => {
  it('Visits meshplay settings page', () => {
    cy.selectProviderNone();
    cy.visit('/settings');
  });
});
