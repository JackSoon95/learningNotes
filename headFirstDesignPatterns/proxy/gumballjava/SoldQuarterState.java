public class SoldQuarterState implements State {
  private static final long serialVersionUID = 2L;

  transient GumballMachine gumballMachine;

  public SoldQuarterState(GumballMachine gumballMachine) {
    this.gumballMachine = gumballMachine;
  }

  public void insertQuarter() {
    System.out.println("Please wait, we're already giving you a gumball");
  }

  public void ejectQuarter() {
    System.out.println("Sorry, you already turned the crank");
  }

  public void turnCrank() {
    System.out.println("Turning twice doesn't get you another gumball!");
  }

  public void dispense() {
    gumballMachine.ReleaseBall();
    if (gumballMachine.getCount() > 0) {
      gumballMachine.setState(gumballMachine.getNoQuarterState());
    } else {
      gumballMachine.setState(gumballMachine.getSoldOutQuarterState());
    }
  }

  public void refill() {
  }
}
