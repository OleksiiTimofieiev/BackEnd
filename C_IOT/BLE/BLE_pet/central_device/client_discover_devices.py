#!/usr/bin/python3
from gi.repository import GLib 
import bluetooth_utils
import bluetooth_constants 
import dbus
import dbus.mainloop.glib 
import sys 

sys.path.insert(0, '.')
adapter_interface = None
mainloop = None
timer_id = None
devices = {}

def interfaces_added(path, interfaces):
    # interfaces is an array of dictionary entries
    if not bluetooth_constants.DEVICE_INTERFACE in interfaces:
        return

    device_properties = interfaces[bluetooth_constants.DEVICE_INTERFACE] 
    if path not in devices:
        print("NEW path :", path) 
        devices[path] = device_properties 
        dev = devices[path]
        if 'Address' in dev:
            print("NEW bdaddr: ", bluetooth_utils.dbus_to_python(dev['Address']))
        if 'Name' in dev:
            print("NEW name  : ",
            bluetooth_utils.dbus_to_python(dev['Name'])) 
        if 'RSSI' in dev:
            print("NEW RSSI  : ",
            bluetooth_utils.dbus_to_python(dev['RSSI'])) 
        print("------------------------------")
def discover_devices(bus):
    global adapter_interface
    global mainloop
    global timer_id
    adapter_path = bluetooth_constants.BLUEZ_NAMESPACE + bluetooth_constants.ADAPTER_NAME
    # acquire an adapter proxy object and its Adapter1 interface so we can call its methods
    adapter_object = bus.get_object(bluetooth_constants.BLUEZ_SERVICE_NAME, adapter_path)
    adapter_interface=dbus.Interface(adapter_object, bluetooth_constants.ADAPTER_INTERFACE)
    # register signal handler functions so we can asynchronously report discovered devices
    # InterfacesAdded signal is emitted by BlueZ when an advertising packet from a device it doesn't
    # already know about is received
    bus.add_signal_receiver(interfaces_added,
        dbus_interface = bluetooth_constants.DBUS_OM_IFACE,
        signal_name = "InterfacesAdded")
    mainloop = GLib.MainLoop()
    adapter_interface.StartDiscovery(byte_arrays=True)
    mainloop.run()
# dbus initialisation steps
dbus.mainloop.glib.DBusGMainLoop(set_as_default=True)
bus = dbus.SystemBus()
print("Scanning") 
discover_devices(bus)