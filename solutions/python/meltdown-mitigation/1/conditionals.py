"""Functions to prevent a nuclear meltdown."""


def is_criticality_balanced(temperature, neutrons_emitted):
    if (temperature < 800 and neutrons_emitted > 500 and (temperature * neutrons_emitted) < 500000):
        return True
    else :
        return False 

def reactor_efficiency(voltage, current, theoretical_max_power):
    generated_power = voltage * current;
    efficiency = generated_power / theoretical_max_power * 100;
    
    if (efficiency >= 80):
       return 'green'
    elif (efficiency >= 60):
        return 'orange'
    elif (efficiency >= 30):
        return 'red'
    else:
        return 'black'


def fail_safe(temperature, neutrons_produced_per_second, threshold):
    criticality_level = (temperature * neutrons_produced_per_second) / threshold * 100

    if (criticality_level < 90):
        return 'LOW'
    elif (90 <= criticality_level <= 110):
        return 'NORMAL'
    else:
        return 'DANGER'